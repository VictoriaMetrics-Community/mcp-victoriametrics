package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/VictoriaMetrics/metricsql"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/tmc/langchaingo/textsplitter"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/resources"
)

var (
	toolExplainQuery = mcp.NewTool("explain_query",
		mcp.WithDescription("Explain how MetricsQL query works"),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Explain Query",
			ReadOnlyHint:    ptr(true),
			DestructiveHint: ptr(false),
			OpenWorldHint:   ptr(true),
		}),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Title("MetricsQL or PromQL expression"),
			mcp.Description(`MetricsQL or PromQL expression for explanation`),
		),
	)
)

func toolExplainQueryHandler(_ context.Context, _ *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := GetToolReqParam[string](tcr, "query", true)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	info, err := getQueryInfo(query)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("error explaining query: %s", err)), nil
	}

	data, err := json.Marshal(info)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("error marshalling query info: %s", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

func getQueryInfo(query string) (map[string]any, error) {
	expr, err := metricsql.Parse(query)
	if err != nil {
		return nil, fmt.Errorf("query parsing error: %w", err)
	}
	types := make(map[string]struct{})
	functions := make(map[string]struct{})
	st := getSyntaxTree(expr, types, functions)
	result := map[string]any{
		"syntax_tree":    st,
		"types_info":     getTypesDescriptions(types),
		"functions_info": getFunctionsInfo(functions),
	}
	return result, nil
}

type FunctionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

var functionsInfo map[string]FunctionInfo

func getFunctionsInfo(functions map[string]struct{}) map[string]FunctionInfo {
	result := make(map[string]FunctionInfo)
	for fn := range functions {
		if info, ok := functionsInfo[fn]; ok {
			result[fn] = info
		} else {
			result[fn] = FunctionInfo{
				Name:        fn,
				Description: fmt.Sprintf("Unknown function %s", fn),
				Category:    "unknown",
			}
		}
	}
	return result
}

func initFunctionsInfo() error {
	var mdSplitter = textsplitter.NewMarkdownTextSplitter(
		textsplitter.WithCodeBlocks(true),
		textsplitter.WithHeadingHierarchy(false),
		textsplitter.WithJoinTableRows(false),
		textsplitter.WithKeepSeparator(false),
		textsplitter.WithReferenceLinks(false),
		textsplitter.WithChunkSize(65536),
		textsplitter.WithChunkOverlap(4096),
	)

	mql, err := resources.DocsDir.ReadFile("vm/docs/victoriametrics/MetricsQL.md")
	if err != nil {
		return fmt.Errorf("error reading MetricsQL documentation: %w", err)
	}

	chunks, err := mdSplitter.SplitText(string(mql))
	if err != nil {
		return fmt.Errorf("error splitting MetricsQL documentation: %w", err)
	}

	functionsInfo = make(map[string]FunctionInfo)
	category := ""
	for _, chunk := range chunks {
		lines := strings.SplitN(chunk, "\n", 2)
		title := lines[0]
		if !strings.HasPrefix(title, "### ") && !strings.HasPrefix(title, "#### ") {
			continue
		}
		if strings.HasPrefix(title, "### ") && strings.Contains(title, "functions") {
			category = strings.TrimSpace(strings.TrimPrefix(title, "### "))
		}
		if category != "" && strings.HasPrefix(title, "#### ") {
			name := strings.TrimSpace(strings.TrimPrefix(title, "#### "))
			content := ""
			if len(lines) > 0 {
				content = lines[1]
			}
			functionsInfo[name] = FunctionInfo{
				Name:        name,
				Description: content,
				Category:    category,
			}
		}
	}

	return nil
}

type TypeDescription struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Fields      map[string]TypeField `json:"fields,omitempty"`
}

type TypeField struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DataType    string `json:"data_type"`
}

var typeDescriptions = map[string]TypeDescription{
	"AggrFuncExpr": {
		Name:        "Aggregate function",
		Description: "AggrFuncExpr represents aggregate function such as `sum(...) by (...)`",
		Fields: map[string]TypeField{
			"name": {
				Description: "Name of the aggregate function, e.g. `sum`, `avg`, etc.",
				DataType:    "string",
			},
			"args": {
				Description: "Arguments of the aggregate function, which can be other expressions.",
				DataType:    "array of Expr",
			},
			"modifier": {
				Description: "Optional modifier for the aggregate function, such as `by (...)` or `without (...)`.",
				DataType:    "ModifierExpr",
			},
			"limit": {
				Description: "Optional limit for the number of output time series. Example: `sum(...) by (...) limit 10` would return maximum 10 time series.",
				DataType:    "int",
			},
		},
	},
	"BinaryOpExpr": {
		Name:        "Binary operation",
		Description: "BinaryOpExpr represents binary operation such as `+`, `-`, `*`, `/`, etc.",
		Fields: map[string]TypeField{
			"op": {
				Description: "Op is the operation itself, i.e. `+`, `-`, `*`, etc.",
				DataType:    "string",
			},
			"bool": {
				Description: "Bool indicates whether `bool` modifier is present. For example, `foo >bool bar`",
				DataType:    "bool",
			},
			"group_modifier": {
				Description: "GroupModifier contains modifier such as \"on\" or \"ignoring\".",
				DataType:    "ModifierExpr",
			},
			"join_modifier": {
				Description: "JoinModifier contains modifier such as \"group_left\" or \"group_right\".",
				DataType:    "ModifierExpr",
			},
			"join_modifier_prefix": {
				Description: "JoinModifierPrefix is an optional prefix to add to labels specified inside group_left() or group_right() lists. The syntax is `group_left(foo,bar) prefix \"abc\"`",
				DataType:    "StringExpr",
			},
			"keep_metric_names": {
				Description: "If KeepMetricNames is set to true, then the operation should keep metric names.",
				DataType:    "bool",
			},
			"left": {
				Description: "Left contains left arg for the `left op right` expression.",
				DataType:    "Expr",
			},
			"right": {
				Description: "Right contains right arg for the `left op right` expression.",
				DataType:    "Expr",
			},
		},
	},
	"DurationExpr": {
		Name:        "Duration",
		Description: "DurationExpr represents a duration, e.g. `5m`, `1h`. Supported suffixes are `s` (seconds), `m` (minutes), `h` (hours), `d` (days), `w` (weeks), and `y` (years).",
		Fields: map[string]TypeField{
			"value": {
				Description: "Value is the duration value as a string, e.g. `5m`, `1h`.",
				DataType:    "string",
			},
		},
	},
	"FuncExpr": {
		Name:        "Function",
		Description: "uncExpr represents MetricsQL function such as `foo(...)`",
		Fields: map[string]TypeField{
			"name": {
				Description: "Name of the function, e.g. `rate`, `histogram_quantile`, etc.",
				DataType:    "string",
			},
			"args": {
				Description: "Arguments of the function, which can be other expressions.",
				DataType:    "array of Expr",
			},
			"keep_metric_names": {
				Description: "If KeepMetricNames is set to true, then the function should keep metric names.",
				DataType:    "bool",
			},
		},
	},
	"LabelFilter": {
		Name:        "Label filter",
		Description: "LabelFilter represents MetricsQL label filter like `foo=\"bar\"`.",
		Fields: map[string]TypeField{
			"label": {
				Description: "Label is the name of the label to filter on.",
				DataType:    "string",
			},
			"value": {
				Description: "Value contains unquoted value for the filter. If IsRegexp is true, then this is a regular expression.",
				DataType:    "string",
			},
			"is_regexp": {
				Description: "IsRegexp represents whether the filter is regesp, i.e. `=~` or `!~`.",
				DataType:    "bool",
			},
			"is_negative": {
				Description: "IsNegative indicates whether the filter is negative, i.e. `!=` or `!~`.",
				DataType:    "bool",
			},
		},
	},
	"MetricExpr": {
		Name: "Metric",
		Description: `MetricExpr represents MetricsQL metric with optional filters, i.e. "foo{...}".
Curly braces may contain or-delimited list of filters. For example:
	x{job="foo",instance="bar" or job="x",instance="baz"}

In this case the filter returns all the series, which match at least one of the following filters:

	x{job="foo",instance="bar"}
	x{job="x",instance="baz"}

This allows using or-delimited list of filters inside rollup functions. For example, the following query calculates rate per each matching series for the given or-delimited filters:

	rate(x{job="foo",instance="bar" or job="x",instance="baz"}[5m])`,
		Fields: map[string]TypeField{
			"label_filters": {
				Description: "LabelFilters is a list of or-delimited label filters. Each filter is an and-delimited list of label filters.",
				DataType:    "array of array of LabelFilter",
			},
		},
	},
	"ModifierExpr": {
		Name:        "Modifier",
		Description: "ModifierExpr represents MetricsQL modifier such as `<op> (...)`",
		Fields: map[string]TypeField{
			"op": {
				Description: "Op is modifier operation.",
				DataType:    "string",
			},
			"args": {
				Description: "Args contains modifier args from parens.",
				DataType:    "array of string",
			},
		},
	},
	"NumberExpr": {
		Name:        "Number",
		Description: "NumberExpr represents a numeric value, e.g. `42`, `3.14`.",
		Fields: map[string]TypeField{
			"value": {
				Description: "N is the parsed number, i.e. `1.23`, `-234`, etc.",
				DataType:    "float64",
			},
		},
	},
	"RollupExpr": {
		Name:        "Rollup",
		Description: "RollupExpr represents MetricsQL expression, which contains at least `offset` or `[...]` part.",
		Fields: map[string]TypeField{
			"expr": {
				Description: "The expression for the rollup. Usually it is MetricExpr, but may be arbitrary expr if subquery is used. https://prometheus.io/blog/2019/01/28/subquery-support/",
				DataType:    "Expr",
			},
			"window": {
				Description: "Window contains optional window value from square brackets. For example, `http_requests_total[5m]` will have Window value `5m`.",
				DataType:    "DurationExpr",
			},
			"step": {
				Description: "Step contains optional step value from square brackets. For example, `foobar[1h:3m]` will have Step value '3m'.",
				DataType:    "DurationExpr",
			},
			"offset": {
				Description: "Offset contains optional value from `offset` part. For example, `foobar{baz=\"aa\"} offset 5m` will have Offset value `5m`.",
				DataType:    "DurationExpr",
			},
			"at": {
				Description: "At contains an optional expression after `@` modifier. For example, `foo @ end()` or `bar[5m] @ 12345`. See https://prometheus.io/docs/prometheus/latest/querying/basics/#modifier",
				DataType:    "Expr",
			},
			"inherit_step": {
				Description: "If set to true, then `foo[1h:]` would print the same",
				DataType:    "bool",
			},
		},
	},
	"StringExpr": {
		Name:        "String",
		Description: "StringExpr represents a string expression, e.g. `\"foo\"`, `\"bar\"`.",
		Fields: map[string]TypeField{
			"value": {
				Description: "Contains unquoted value for string expression.",
				DataType:    "string",
			},
		},
	},
}

func getTypesDescriptions(ts map[string]struct{}) map[string]TypeDescription {
	result := make(map[string]TypeDescription)
	for t := range ts {
		if desc, ok := typeDescriptions[t]; ok {
			result[t] = desc
		} else {
			result[t] = TypeDescription{
				Name:        t,
				Description: fmt.Sprintf("Unknown type %s", t),
			}
		}
	}
	return result
}

func getSyntaxTree(
	e metricsql.Expr,
	types map[string]struct{},
	functions map[string]struct{},
) map[string]any {
	if e == nil {
		return nil
	}
	result := make(map[string]any)
	switch n := e.(type) {
	case *metricsql.AggrFuncExpr:
		types["AggrFuncExpr"] = struct{}{}
		result["type"] = "AggrFuncExpr"
		functions[n.Name] = struct{}{}
		result["name"] = n.Name
		result["limit"] = n.Limit
		args := make([]any, 0)
		for _, arg := range n.Args {
			argInfo := getSyntaxTree(arg, types, functions)
			args = append(args, argInfo)
		}
		result["args"] = args
		result["modifier"] = getSyntaxTree(&n.Modifier, types, functions)
	case *metricsql.BinaryOpExpr:
		types["BinaryOpExpr"] = struct{}{}
		result["type"] = "BinaryOpExpr"
		result["op"] = n.Op
		result["bool"] = n.Bool
		result["group_modifier"] = getSyntaxTree(&n.GroupModifier, types, functions)
		result["join_modifier"] = getSyntaxTree(&n.JoinModifier, types, functions)
		result["join_modifier_prefix"] = n.JoinModifierPrefix
		result["left"] = getSyntaxTree(n.Left, types, functions)
		result["right"] = getSyntaxTree(n.Right, types, functions)
		result["keep_metric_name"] = n.KeepMetricNames
	case *metricsql.DurationExpr:
		types["DurationExpr"] = struct{}{}
		result["type"] = "DurationExpr"
		result["value"] = n.AppendString(nil)
	case *metricsql.FuncExpr:
		types["FuncExpr"] = struct{}{}
		result["type"] = "FuncExpr"
		functions[n.Name] = struct{}{}
		result["name"] = n.Name
		args := make([]any, 0)
		for _, arg := range n.Args {
			argInfo := getSyntaxTree(arg, types, functions)
			args = append(args, argInfo)
		}
		result["args"] = args
		result["keep_metric_name"] = n.KeepMetricNames
	case *metricsql.LabelFilter:
		types["LabelFilter"] = struct{}{}
		result["type"] = "LabelFilter"
		result["label"] = n.Label
		result["value"] = n.Value
		result["is_regexp"] = n.IsRegexp
		result["is_negative"] = n.IsNegative
	case *metricsql.MetricExpr:
		types["MetricExpr"] = struct{}{}
		result["type"] = "MetricExpr"
		labelFilterss := make([]any, 0)
		for _, labelFilters := range n.LabelFilterss {
			fs := make([]any, 0)
			for _, filter := range labelFilters {
				fsInfo := getSyntaxTree(&filter, types, functions)
				fs = append(fs, fsInfo)
			}
			labelFilterss = append(labelFilterss, fs)
		}
		result["label_filters"] = labelFilterss
	case *metricsql.ModifierExpr:
		types["ModifierExpr"] = struct{}{}
		result["type"] = "ModifierExpr"
		result["op"] = n.Op
		result["args"] = n.Args
	case *metricsql.NumberExpr:
		types["NumberExpr"] = struct{}{}
		result["type"] = "NumberExpr"
		result["value"] = n.N
	case *metricsql.RollupExpr:
		types["RollupExpr"] = struct{}{}
		result["type"] = "RollupExpr"
		result["expr"] = getSyntaxTree(n.Expr, types, functions)
		if n.Window != nil {
			result["window"] = getSyntaxTree(n.Window, types, functions)
		}
		if n.Step != nil {
			result["step"] = getSyntaxTree(n.Step, types, functions)
		}
		if n.Offset != nil {
			result["offset"] = getSyntaxTree(n.Offset, types, functions)
		}
		if n.At != nil {
			result["at"] = getSyntaxTree(n.At, types, functions)
		}
		result["inherit_step"] = n.InheritStep
	case *metricsql.StringExpr:
		types["StringExpr"] = struct{}{}
		result["type"] = "StringExpr"
		result["value"] = n.S
	}
	return result
}

func RegisterToolExplainQuery(s *server.MCPServer, c *config.Config) {
	err := initFunctionsInfo()
	if err != nil {
		panic(fmt.Sprintf("error initializing functions info: %s", err))
	}
	s.AddTool(toolExplainQuery, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolExplainQueryHandler(ctx, c, request)
	})
}
