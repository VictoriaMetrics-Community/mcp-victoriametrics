---
title : "AI Assistants"
menu:
  docs:
    parent: "integrations"
    identifier: "cloud-integrations-ai-assistants"
---

VictoriaMetrics Cloud exposes a hosted [MCP (Model Context Protocol) server](https://github.com/VictoriaMetrics/mcp-victoriametrics)
that allows MCP-compatible AI assistants to query and explore your metrics deployments using natural language.

This integration enables AI tools to access your VictoriaMetrics Cloud metrics directly, making it
easier to investigate incidents, explore data, and build queries without leaving your AI assistant.

> Only the **metrics** MCP server is hosted by VictoriaMetrics Cloud today. Hosted MCP servers for
> VictoriaLogs and VictoriaTraces are not available yet. See [Using AI Assistants with VictoriaLogs](https://docs.victoriametrics.com/victoriametrics-cloud/integrations/ai-assistants/#using-ai-assistants-with-victorialogs)
> below for how to run it yourself in the meantime.

## Integrating AI Assistants via MCP

Unlike other VictoriaMetrics Cloud integrations, the hosted MCP server uses an **API key** for
authentication rather than a per-deployment access token. You can configure the deployments this
API key has access to.

You can manage them in the [API keys](https://console.victoriametrics.cloud/organization/api_keys)
page in the console, or read more in the dedicated [documentation page](https://docs.victoriametrics.com/victoriametrics-cloud/api/).

> [!WARNING]
> Avoid sharing configuration files that contain your API key in plaintext. If your AI tool supports
> environment variables in MCP configurations, use them instead of storing the key directly.
> For example, `${VMCLOUD_API_KEY}`.

### Connecting your AI Assistant to the hosted MCP server
To connect your AI assistant to VictoriaMetrics Cloud via MCP, visit the [cloud console](https://console.victoriametrics.cloud/integrations/ai-assistants),
or follow this interactive guide:

{{< cloud-console-iframe "ai-assistants" >}}

### Running the MCP Server Locally

If you prefer to run the MCP server on your own infrastructure, you can set it up locally by
providing the bearer token for your deployment, its instance type, and entry point. See the
[setup instructions](https://github.com/VictoriaMetrics/mcp-victoriametrics#setup-in-clients) in
the mcp-victoriametrics repository for details.

## Using AI Assistants with VictoriaLogs

Unlike metrics, VictoriaMetrics Cloud does not (yet) host an MCP server for VictoriaLogs. You can
still connect AI assistants to your VictoriaLogs deployment by running the
[mcp-victorialogs](https://github.com/VictoriaMetrics/mcp-victorialogs) server yourself.

All you need is:
- The **bearer token** ([access token](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/)) for your VictoriaLogs deployment.
- The deployment's **entry point URL**, bearing in mind that no specific query path is required since
the MCP server handles routing to the right VictoriaLogs API endpoints internally.

Because AI assistants and their MCP configuration formats vary widely, refer to the
[mcp-victorialogs repository](https://github.com/VictoriaMetrics/mcp-victorialogs) for setup
instructions and configuration examples matching your assistant of choice.

## Using AI Assistants with VictoriaTraces

There is no hosted or dedicated MCP server for VictoriaTraces yet. Since VictoriaTraces stores trace
spans as structured data queryable with [LogsQL](https://docs.victoriametrics.com/victorialogs/logsql/),
the [mcp-victorialogs](https://github.com/VictoriaMetrics/mcp-victorialogs) server can be pointed at
a VictoriaTraces deployment entry point to explore spans, using a read
[access token](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/) as
bearer token.

To learn more about AI capabilities across VictoriaMetrics products, see the
[AI Tools](https://docs.victoriametrics.com/ai-tools/) documentation.
