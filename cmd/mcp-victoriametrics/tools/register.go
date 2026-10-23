package tools

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

func Register(s *server.MCPServer, c *config.Config) {
	registerToolQuery(s, c)
	registerToolFlags(s, c)
	registerToolRules(s, c)
	registerToolAlerts(s, c)
	registerToolLabels(s, c)
	registerToolSeries(s, c)
	registerToolExport(s, c)
	registerToolTenants(s, c)
	registerToolMetrics(s, c)
	registerToolTestRules(s, c)
	registerToolTSDBStatus(s, c)
	registerToolQueryRange(s, c)
	registerToolTopQueries(s, c)
	registerToolMetricStats(s, c)
	registerToolLabelValues(s, c)
	registerToolExplainQuery(s, c)
	registerToolActiveQueries(s, c)
	registerToolDocumentation(s, c)
	registerToolPrettifyQuery(s, c)
	registerToolMetricsMetadata(s, c)
	registerToolMetricRelabelDebug(s, c)
	registerToolRetentionFiltersDebug(s, c)
	registerToolDownsamplingFiltersDebug(s, c)
	registerToolTiers(s, c)
	registerToolRegions(s, c)
	registerToolRuleFile(s, c)
	registerToolDeployments(s, c)
	registerToolAccessTokens(s, c)
	registerToolRuleFilenames(s, c)
	registerToolCloudProviders(s, c)
}
