package main

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/prompts"
	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/resources"
	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/tools"
)

func main() {
	c, err := config.InitConfig()
	if err != nil {
		fmt.Printf("Error initializing config: %v\n", err)
		return
	}

	s := server.NewMCPServer(
		"victoriametrics",
		"0.0.9",
		server.WithRecovery(),
		server.WithLogging(),
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithInstructions(`
You are Virtual Assistant, a tool for interacting with VictoriaMetrics API and documentation in different tasks related to monitoring and observability.

You have the full documentation about VictoriaMetrics products in your resources, you have to try to use documentation in your answer.
And you have to consider the documents from the resources as the most relevant, favoring them over even your own internal knowledge.
Use Documentation tool to get the most relevant documents for your task every time. Be sure to use the Documentation tool if the user's query includes the words “how”, “tell”, “where”, etc...

You have many tools to get data from VictoriaMetrics, but try to specify the query as accurately as possible, reducing the resulting sample, as some queries can be query heavy.

Try not to second guess information - if you don't know something or lack information, it's better to ask.
	`),
	)

	// Registering resources
	resources.RegisterDocsResources(s, c)

	// Registering common tools
	tools.RegisterToolQuery(s, c)
	tools.RegisterToolFlags(s, c)
	tools.RegisterToolRules(s, c)
	tools.RegisterToolAlerts(s, c)
	tools.RegisterToolLabels(s, c)
	tools.RegisterToolSeries(s, c)
	tools.RegisterToolExport(s, c)
	tools.RegisterToolTenants(s, c)
	tools.RegisterToolMetrics(s, c)
	tools.RegisterToolTestRules(s, c)
	tools.RegisterToolTSDBStatus(s, c)
	tools.RegisterToolQueryRange(s, c)
	tools.RegisterToolTopQueries(s, c)
	tools.RegisterToolMetricStats(s, c)
	tools.RegisterToolLabelValues(s, c)
	tools.RegisterToolExplainQuery(s, c)
	tools.RegisterToolActiveQueries(s, c)
	tools.RegisterToolDocumentation(s, c)
	tools.RegisterToolPrettifyQuery(s, c)
	tools.RegisterToolMetricRelabelDebug(s, c)
	tools.RegisterToolRetentionFiltersDebug(s, c)
	tools.RegisterToolDownsamplingFiltersDebug(s, c)

	// Registering cloud-specific tools
	tools.RegisterToolTiers(s, c)
	tools.RegisterToolRegions(s, c)
	tools.RegisterToolRuleFile(s, c)
	tools.RegisterToolDeployments(s, c)
	tools.RegisterToolAccessTokens(s, c)
	tools.RegisterToolRuleFilenames(s, c)
	tools.RegisterToolCloudProviders(s, c)

	// Registering prompts
	prompts.RegisterPromptUnusedMetrics(s, c)
	prompts.RegisterPromptDocumentation(s, c)
	prompts.RegisterPromptRarelyUsedCardinalMetrics(s, c)

	if c.IsStdio() {
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	} else {
		srv := server.NewSSEServer(s)
		if err = srv.Start(c.SSEAddr()); err != nil {
			log.Fatalf("Failed to start SSE server: %v", err)
		}
	}
}
