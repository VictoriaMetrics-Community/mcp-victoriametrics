# VictoriaMetrics MCP Server

[![Latest Release](https://img.shields.io/github/v/release/VictoriaMetrics-Community/mcp-victoriametrics?sort=semver&label=&filter=!*-victorialogs&logo=github&labelColor=gray&color=gray&link=https%3A%2F%2Fgithub.com%2FVictoriaMetrics-Community%2Fmcp-victoriametrics%2Freleases%2Flatest)](https://github.com/VictoriaMetrics-Community/mcp-victoriametrics/releases)
[![smithery badge](https://smithery.ai/badge/@VictoriaMetrics-Community/mcp-victoriametrics)](https://smithery.ai/server/@VictoriaMetrics-Community/mcp-victoriametrics)
![License](https://img.shields.io/github/license/VictoriaMetrics-Community/mcp-victoriametrics?labelColor=green&label=&link=https%3A%2F%2Fgithub.com%2FVictoriaMetrics-Community%2Fmcp-victoriametrics%2Fblob%2Fmain%2FLICENSE)
![Slack](https://img.shields.io/badge/Join-4A154B?logo=slack&link=https%3A%2F%2Fslack.victoriametrics.com)
![X](https://img.shields.io/twitter/follow/VictoriaMetrics?style=flat&label=Follow&color=black&logo=x&labelColor=black&link=https%3A%2F%2Fx.com%2FVictoriaMetrics)
![Reddit](https://img.shields.io/reddit/subreddit-subscribers/VictoriaMetrics?style=flat&label=Join&labelColor=red&logoColor=white&logo=reddit&link=https%3A%2F%2Fwww.reddit.com%2Fr%2FVictoriaMetrics)

The implementation of [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server for [VictoriaMetrics](https://docs.victoriametrics.com/victoriametrics/).

This provides access to your VictoriaMetrics instance and seamless integration with [VictoriaMetrics APIs](https://docs.victoriametrics.com/victoriametrics/url-examples/) and [documentation](https://docs.victoriametrics.com/).
It can give you a comprehensive interface for monitoring, observability, and debugging tasks related to your VictoriaMetrics instances, enable advanced automation and interaction capabilities for engineers and tools.

## Features

This MCP server allows you to use almost all read-only APIs of VictoriaMetrics, i.e. all functions available in [VMUI](https://docs.victoriametrics.com/#vmui):

- Querying metrics and exploring data (even drawing graphs if your client supports it)
- Listing and exporting available metrics, labels, labels values and entire series
- Analyzing your alerting and recording rules and alerts
- Showing parameters of your VictoriaMetrics instance
- Exploring cardinality of your data and metrics usage statistics
- Analyzing, tracing, prettifying and explaining your queries
- Debugging your relabeling rules, downsampling and retention policy configurations 
 
In addition, the MCP server contains embedded up-to-date documentation and is able to search it without online access.

More details about the exact available tools and prompts can be found in the [Usage](#usage) section.

You can combine functionality of tools, docs search in your prompts and invent great usage scenarios for your VictoriaMetrics instance.
Just check the [Dialog example](#dialog-example) section to see how it can work.
And please note the fact that the quality of the MCP Server and its responses depends very much on the capabilities of your client and the quality of the model you are using.

You can also combine the MCP server with other observability or doc search related MCP Servers and get even more powerful results.

## Requirements

- [VictoriaMetrics](https://docs.victoriametrics.com/victoriametrics/) instance ([single-node](https://docs.victoriametrics.com/victoriametrics/single-server-victoriametrics/) or [cluster](https://docs.victoriametrics.com/victoriametrics/cluster-victoriametrics/))
- Go 1.24 or higher (if you want to build from source)

## Installation

### Go

```bash
go install github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics@latest
```

### Source Code

```bash
git clone https://github.com/VictoriaMetrics-Community/mcp-victoriametrics.git
cd mcp-victoriametrics
go build -o bin/mcp-victoriametrics ./cmd/mcp-victoriametrics/main.go

# after that add bin/mcp-victoriametrics file to your PATH
```

### Binaries

Just download the latest release from [Releases](https://github.com/VictoriaMetrics-Community/mcp-victoriametrics/releases) page and put it to your PATH.

### Docker

Coming soon...

### Smithery

To install VictoriaMetrics MCP Server for your client automatically via Smithery, yo can use the following commands:

```bash
# Get the list of supported MCP clients
npx -y @smithery/cli list clients
#Available clients:
#  claude
#  cline
#  windsurf
#  roocode
#  witsy
#  enconvo
#  cursor
#  vscode
#  vscode-insiders
#  boltai
#  amazon-bedrock

# Install VictoriaMetrics MCP server for your client
npx -y @smithery/cli install @VictoriaMetrics-Community/mcp-victoriametrics --client <YOUR-CLIENT-NAME>
# and follow the instructions
```

## Configuration

MCP Server for VictoriaMetrics is configured via environment variables:

| Variable | Description | Required | Default | Allowed values |
|----------|-------------|----------|---------|---------|
| `VM_INSTANCE_ENTRYPOINT` | URL to VictoriaMetrics instance | Yes | - | - |
| `VM_INSTANCE_TYPE` | Type of VictoriaMetrics instance | Yes | - | `single`, `cluster` |
| `VM_INSTANCE_BEARER_TOKEN` | Authentication token for VictoriaMetrics API | No | - | - |
| `MCP_SERVER_MODE` | Server operation mode | No | `stdio` | `stdio`, `sse` |
| `MCP_SSE_ADDR` | Address for SSE server to listen on | No | `localhost:8080` | - |

### Сonfiguration examples

```bash
# For a single-node instance
export VM_INSTANCE_ENTRYPOINT="http://localhost:8428"
export VM_INSTANCE_TYPE="single"
export VM_INSTANCE_BEARER_TOKEN="your-token"

# For a cluster
export VM_INSTANCE_ENTRYPOINT="https://play.victoriametrics.com"
export VM_INSTANCE_TYPE="cluster"

# Server mode
export MCP_SERVER_MODE="sse"
export MCP_SSE_ADDR="0.0.0.0:8080"
```

## Setup in clients

### Cursor

Go to: `Settings` -> `Cursor Settings` -> `MCP` -> `Add new global MCP server` and paste the following configuration into your Cursor `~/.cursor/mcp.json` file:

```json
{
  "mcpServers": {
    "victoriametrics": {
      "command": "/path/to/mcp-victoriametrics",
      "env": {
        "VM_INSTANCE_ENTRYPOINT": "<YOUR_VM_INSTANCE>",
        "VM_INSTANCE_TYPE": "<YOUR_VM_INSTANCE_TYPE>",
        "VM_INSTANCE_BEARER_TOKEN": "<YOUR_VM_BEARER_TOKEN>"
      }
    }
  }
}
```

See [Cursor MCP docs](https://docs.cursor.com/context/model-context-protocol) for more info.

### Claude Desktop

Add this to your Claude Desktop `claude_desktop_config.json` file (you can find it if open `Settings` -> `Developer` -> `Edit config`):

```json
{
  "mcpServers": {
    "victoriametrics": {
      "command": "/path/to/mcp-victoriametrics",
      "env": {
        "VM_INSTANCE_ENTRYPOINT": "<YOUR_VM_INSTANCE>",
        "VM_INSTANCE_TYPE": "<YOUR_VM_INSTANCE_TYPE>",
        "VM_INSTANCE_BEARER_TOKEN": "<YOUR_VM_BEARER_TOKEN>"
      }
    }
  }
}
```

See [Claude Desktop MCP docs](https://modelcontextprotocol.io/quickstart/user) for more info.

### Claude Code

Run the command:

```sh
claude mcp add victoriametrics -- /path/to/mcp-victoriametrics \
  -e VM_INSTANCE_ENTRYPOINT=<YOUR_VM_INSTANCE> \
  -e VM_INSTANCE_TYPE=<YOUR_VM_INSTANCE_TYPE>
  -e VM_INSTANCE_BEARER_TOKEN=<YOUR_VM_BEARER_TOKEN>
```

See [Claude Code MCP docs](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/tutorials#set-up-model-context-protocol-mcp) for more info.

### Visual Studio Code

Add this to your VS Code MCP config file:

```json
{
  "servers": {
    "victoriametrics": {
      "type": "stdio",
      "command": "/path/to/mcp-victoriametrics",
      "env": {
        "VM_INSTANCE_ENTRYPOINT": "<YOUR_VM_INSTANCE>",
        "VM_INSTANCE_TYPE": "<YOUR_VM_INSTANCE_TYPE>",
        "VM_INSTANCE_BEARER_TOKEN": "<YOUR_VM_BEARER_TOKEN>"
      }
    }
  }
}
```

See [VS Code MCP docs](https://code.visualstudio.com/docs/copilot/chat/mcp-servers) for more info.

### Zed

Add the following to your Zed config file:

```json
  "context_servers": {
    "victoriametrics": {
      "command": {
        "path": "/path/to/mcp-victoriametrics",
        "args": [],
        "env": {
          "VM_INSTANCE_ENTRYPOINT": "<YOUR_VM_INSTANCE>",
          "VM_INSTANCE_TYPE": "<YOUR_VM_INSTANCE_TYPE>",
          "VM_INSTANCE_BEARER_TOKEN": "<YOUR_VM_BEARER_TOKEN>"
        }
      },
      "settings": {}
    }
  }
```

See [Zed MCP docs](https://zed.dev/docs/ai/mcp) for more info.

### JetBrains IDEs

- Open `Settings` -> `Tools` -> `AI Assistant` -> `Model Context Protocol (MCP)`.
- Click `Add (+)`
- Select `As JSON`
- Put the following to the input field:

```json
{
  "mcpServers": {
    "victoriametrics": {
      "command": "/path/to/mcp-victoriametrics",
      "env": {
        "VM_INSTANCE_ENTRYPOINT": "<YOUR_VM_INSTANCE>",
        "VM_INSTANCE_TYPE": "<YOUR_VM_INSTANCE_TYPE>",
        "VM_INSTANCE_BEARER_TOKEN": "<YOUR_VM_BEARER_TOKEN>"
      }
    }
  }
}
```

### Windsurf

Add the following to your Windsurf MCP config file.

```json
{
  "mcpServers": {
    "victoriametrics": {
      "command": "/path/to/mcp-victoriametrics",
      "env": {
        "VM_INSTANCE_ENTRYPOINT": "<YOUR_VM_INSTANCE>",
        "VM_INSTANCE_TYPE": "<YOUR_VM_INSTANCE_TYPE>",
        "VM_INSTANCE_BEARER_TOKEN": "<YOUR_VM_BEARER_TOKEN>"
      }
    }
  }
}
```

See [Windsurf MCP docs](https://docs.windsurf.com/windsurf/mcp) for more info.

### Amazon Bedrock

Coming soon....

### Using Docker instead of binary

Coming soon...

## Usage

After [installing](#installation) and [configuring](#setup-in-clients) the MCP server, you can start using it with your favorite MCP client.

You can start dialog with AI assistant from the phrase:

```
Use MCP VictoriaMetrics in the following answers
```

But it's not required, you can just start asking questions and the assistant will automatically use the tools and documentation to provide you with the best answers.
Just take a look into [Dialog example](#dialog-example) section for better understanding what you can do with it.

### Toolset

MCP VictoriaMetrics provides numerous tools for interacting with your VictoriaMetrics instance.

Here's a list of available tools:

| Tool | Description                                               |
|------|-----------------------------------------------------------|
| `query` | Execute instant PromQL/MetricsQL queries                  |
| `query_range` | Execute range PromQL/MetricsQL queries over a time period |
| `metrics` | List available metrics                                    |
| `labels` | List available label names                                |
| `label_values` | List values for a specific label                          |
| `series` | List available time series                                |
| `export` | Export raw time series data to JSON or CSV                |
| `rules` | View alerting and recording rules                         |
| `alerts` | View current alerts (firing and pending)                  |
| `flags` | View non-default flags of the VictoriaMetrics instance    |
| `metric_statistics` | Get metrics usage (in queries) statistics                 |
| `active_queries` | View currently executing queries                          |
| `top_queries` | View most frequent or slowest queries                     |
| `tsdb_status` | View TSDB cardinality statistics                          |
| `tenants` | List available tenants in multi-tenant cluster setup      |
| `documentation` | Search in embedded VictoriaMetrics documentation          |
| `metric_relabel_debug` | Debug Prometheus-compatible relabeling rules              |
| `downsampling_filters_debug` | Debug downsampling configuration                          |
| `retention_filters_debug` | Debug retention filters configuration                     |
| `prettify_query` | Prettify and format PromQL/MetricsQL queries              |
| `explain_query` | Parse PromQL/MetricsQL queries and explain how it works   |

### Prompts

The server includes pre-defined prompts for common tasks.

These are just examples at the moment, the prompt library will be added to in the future:

| Prompt | Description |
|--------|-------------|
| `unused_metrics` | Find metrics that are never queried and create relabel configs to stop ingesting them |
| `documentation` | Search VictoriaMetrics documentation for specific topics |
| `rarely_used_metrics_with_high_cardinality` | Identify high-cardinality metrics that are rarely queried |

## Dialog example

This is just example of dialog with Assistant using MCP Server that can help to understand how it can work (but's it's real dialog with using the Victoria Metrics MCP server with Claude Desktop and configured to [public playground instance](https://play.victoriametrics.com)):

- > **You**: What tenants are available in my VM?
- > **Assistant** *In your VictoriaMetrics instance, the following tenants are available `0:0` (default tenant), `1:0`, `10:0`, `100:0`*
- > **You**: From now on, please use only default tenant
- > **Assistant**: *I understand. From now on, I'll use only the default tenant (0:0) for all operations with your VictoriaMetrics instance.*
- > **You**: Which metrics do we have?
- > **Assistant**: *There are many metrics available in your VictoriaMetrics instance. The list includes metrics from various systems and components, such as:*
  >
  > - *System metrics (`node_`, `machine_`, etc.),*
  > - *Kubernetes metrics (`kube_`, `container_`, `apiserver_`, etc.),*
  > - *VictoriaMetrics-specific metrics (`vm_`, `vmagent_`, `vmalert_`, etc.),*
  > - *Go runtime metrics (`go_`),*
  > - *Prometheus-related metrics (`scrape_`, `promhttp_`),*
  > - *Various application metrics (`argocd_`, `alertmanager_`).*
  >
  > *This is a comprehensive list of over 2000 metrics that are available for querying. Is there a specific type of metric you're interested in exploring further?*
- > **You**: Ok, let's work with kubernetes metrics from now on. Please show me which of these metrics never queried?
- > **Assistant**: *Based on the results, I can see that there are two Kubernetes metrics that have never been queried (query count is 0):*
  >
  > - *`kube_persistentvolumeclaim_deletion_timestamp` - This metric provides information about when PersistentVolumeClaims were marked for deletion.*
  > - *`kube_pod_init_container_status_last_terminated_reason` - This metric shows the reason why init containers in pods were last terminated.*
  >
  > *These metrics might be valuable for monitoring and troubleshooting purposes:*
  >
  > - *The PVC deletion timestamp metric could help track deletion patterns or identify stuck PVCs.*
  > - *The init container termination reason metric could help identify common failure patterns in your initialization containers.*
  >
  > *Would you like me to explain more about what these metrics track or suggest some useful queries to start monitoring them?*
- > **You**: Yes, please. And also, can you create relabeling rules to stop ingesting these metrics?
- > **Assistant**: *I'll create a relabel configuration to exclude these never-queried metrics from collection. This will help reduce storage and processing overhead. Here's the relabel configuration you can add to your VictoriaMetrics scrape configuration:*
  >
  > ```yaml
  > relabel_configs:
  > - action: drop
  >   source_labels: [__name__]
  >   regex: kube_persistentvolumeclaim_deletion_timestamp|kube_pod_init_container_status_last_terminated_reason
  > ```
  >
  > *This configuration will:*
  > 1. *Look at each metric's name (the `__name__` label)*
  > 2. *Check if it matches either of the specified metrics*
  > 3. *Drop the metrics that match, preventing them from being collected and stored*
  >
  > *You can add this to your VictoriaMetrics configuration file or to the relevant scrape job configurations in your Prometheus configuration if you're using Prometheus as a scraper for VictoriaMetrics.*
  > *If you want to test this relabel configuration before implementing it, you can use the metric relabel debug tool to see how it would affect your metrics.*
- > **You**: Yes, please debug this relabel config on real time series to see how it would affect my metrics
- > **Assistant**: *Perfect! The relabel config works as expected:*
  >
  > - *The first two test cases show that both `kube_persistentvolumeclaim_deletion_timestamp` and `kube_pod_init_container_status_last_terminated_reason` metrics are correctly dropped (resulting labels are empty `{}`).*
  > - *The third test case confirms that other metrics like `kube_pod_status_phase` are preserved and not affected by our drop rule.*
  >
  > *You can confidently add this relabel configuration to your VictoriaMetrics or Prometheus configuration to stop collecting these unused metrics. This will save storage space and reduce processing overhead without affecting your current monitoring capabilities.*
- > **You**: But i'm using vmagent resource of VictoriaMetrics operator, could you please give me part of specification for my agent with this relabel config?
- > **Assistant**: *You're right, since you're using the VictoriaMetrics Operator and the VMAgent custom resource, the configuration needs to be in the appropriate format for that CR. Here's how you can add the relabel configuration to your VMAgent specification:*
  >
  > ```yaml
  > apiVersion: operator.victoriametrics.com/v1beta1
  > kind: VMAgent
  > metadata:
  > name: your-vmagent-name
  > spec:
  >   # ... your existing configuration ...
  >   
  >   # Add this section to configure global metric relabeling
  >   inlineRelabelConfig:
  >     - action: drop
  >       sourceLabels: ["__name__"]
  >       regex: kube_persistentvolumeclaim_deletion_timestamp|kube_pod_init_container_status_last_terminated_reason
  >   
  >   # ... rest of your configuration ...
  > ```
  >
  > *Key points:*
  >
  > - *In the VMAgent spec, you use `inlineRelabelConfig` for inline relabeling rules*
  > - *The field is called `sourceLabels` (not `source_labels` as in Prometheus config)*
  > - *This is a global relabeling that will be applied to all metrics before they're stored*
- > **You**: I decided not to remove those metrics, but just lower the sample rate for them using Downsampling. What downsampling parameters i have now?
- > **Assistant**: *I can see that you currently have the following downsampling configuration:*
  >
  > ```
  > -downsampling.period="7d:1m,30d:5m,180d:30m"
  > ```
  >
  > *This means:*
  >
  > - *Data newer than 7 days is kept at its original resolution*
  > - *Data between 7-30 days old is downsampled to 1-minute resolution*
  > - *Data between 30-180 days old is downsampled to 5-minute resolution*
  > - *Data older than 180 days is downsampled to 30-minute resolution*
  >
  > *Now, if you want to add specific downsampling rules for those rarely-queried metrics, you can use the downsampling filters debug tool to test how new rules would work. Let's test a configuration where we apply more aggressive downsampling to those specific metrics:*
  >
  > *...getting series and using downsampling filter debugger...*
  >
  > *Great! Now our test is working correctly. Based on the current configuration and our testing, here's a suggested downsampling configuration for your VMAgent CR that would apply more aggressive downsampling to those rarely-queried metrics:*
  >
  > ```yaml
  > apiVersion: operator.victoriametrics.com/v1beta1
  > kind: VMCluster
  > metadata:
  >   name: monitoring-vm-cluster
  > spec:
  >   # ... your existing configuration ...
  >   
  >   # Update the vmstorage section with the new downsampling parameters
  >   vmstorage:
  >     extraArgs:
  >       downsampling.period: '{__name__=\"kube_persistentvolumeclaim_deletion_timestamp\"}:1d:10m,{__name__=\"kube_pod_init_container_status_last_terminated_reason\"}:1d:10m,7d:1m,30d:5m,180d:30m'
  >   
  >   # ... rest of your configuration ...
  > ``` 
  >
  > *This configuration:*
  >
  > *- Applies 10-minute resolution downsampling for the rarely-queried metrics after just 1 day*
  > *- Maintains your existing downsampling policy for all other metrics*
  >
  > *The order of the downsampling rules is important - more specific rules should be listed first.*

During this dialog, the assistant was using the following tools:

- `tenants` to get the list of available tenants
- `documentation` to get information about functionality and data formats
- `metrics` to get the list of available metrics
- `metrics_staistics` to get the information about metrics usage
- `series` to get the time series for debugging
- `metric_relabel_debug` to debug relabeling rules
- `flags` to get the information about instance parameters
- `downsampling_filters_debug` to debug downsampling configuration

But you can use any other tools and combine them in your own way.

## Roadmap

- [x] Support "Prettify query" tool (done in [`v0.0.5`](https://github.com/VictoriaMetrics-Community/mcp-victoriametrics/releases/tag/v0.0.65))
- [x] Support "Explain query" tool (done in [`v0.0.6`](https://github.com/VictoriaMetrics-Community/mcp-victoriametrics/releases/tag/v0.0.6))
- [ ] Support tool for analysis of [Query execution statistics](https://docs.victoriametrics.com/victoriametrics/query-stats/)
- [ ] Suppport tool for [unit-testing of alerting and recording rules](https://docs.victoriametrics.com/victoriametrics/vmalert-tool/)
- [ ] Support optional integration with [VictoriaMetrics Cloud](https://victoriametrics.com/products/cloud/) (via [API keys](https://docs.victoriametrics.com/victoriametrics-cloud/api/))
- [ ] Add some extra knowledge to server in addition to current documentation tool:
  - [ ] [VictoriaMetrics blog](https://victoriametrics.com/blog/) posts
  - [ ] Github issues
  - [ ] Public slack chat history
  - [ ] CRD schemas
  - [ ] Alerting and recording rule sets
- [ ] Implement multitenant version of MCP (that will support several deployments)
- [ ] Add flags/configs validation tool
- [ ] Support tools for vmagent API

## Disclaimer

AI services and agents along with MCP servers like this cannot guarantee the accuracy, completeness and reliability of results.
You should double check the results obtained with AI.
The quality of the MCP Server and its responses depends very much on the capabilities of your client and the quality of the model you are using.

## Contributing

Contributions to the MCP VictoriaMetrics project are welcome! Please feel free to submit issues, feature requests, or pull requests.
