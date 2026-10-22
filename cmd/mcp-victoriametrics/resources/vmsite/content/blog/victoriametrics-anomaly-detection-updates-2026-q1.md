---
draft: false
page: blog blog_post
authors:
  - Fred Navruzov
date: 2026-04-23
title: "What's new in VictoriaMetrics Anomaly Detection (Q1 2026)"
summary: "Q1 2026 brought incremental but important updates to VictoriaMetrics Anomaly Detection: UI improvements, AI assistance inside the UI, a public traces playground, new false-positive reduction controls, and continued resource optimizations."
description: "What's new in in the latest VictoriaMetrics Anomaly Detection Updates, from AI-assisted anomaly analysis to UI improvements and better backtesting workflows."
enableComments: false
featuredPost: false
categories:
  - Observability
  - Product News
tags:
  - anomaly detection
  - vmanomaly
  - observability
  - victoriametrics

images:
  - /blog/victoriametrics-anomaly-detection-updates-2026-q1/preview.webp
---

> **Summary**: In this Q1 2026 update, VictoriaMetrics Anomaly Detection and vmanomaly evolved from a set of powerful capabilities into a smoother end-to-end workflow for real users. The quarter included 8 service releases, 7 anomaly detection UI iterations, AI copilot integration inside the UI, a new public playground for traces, better visualization for business arguments, improved support for reducing false positives, and a broad set of bug fixes and resource optimizations. We also highlight documentation and configuration improvements that make vmanomaly easier to adapt for internal teams, plus upcoming work on agentic skills, online models, root-cause analysis, dynamic presets, and VictoriaMetrics Cloud integration. Overall, these anomaly detection updates are focused on making it easier to try, tune, operate, and improve anomaly detection across metrics, logs, and traces.

Following our [2025 updates](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2025/), here we recap how VictoriaMetrics Anomaly Detection evolved in Q1 2026.

Stay tuned for upcoming content on [anomaly detection](https://victoriametrics.com/tags/anomaly-detection/).

**Series posts:**

- [H1'24 Updates](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2024-q1-2/)
- [H2'24 Updates](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2024-q3/)
- [2025 Updates](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2025/)
- **Q1'26 Updates - you are here**

### Introduction

> [!NOTE] 
For those who might not be familiar: [VictoriaMetrics Anomaly Detection (vmanomaly)](https://victoriametrics.com/products/enterprise/anomaly-detection/) is our anomaly detection product available to enterprise customers. It helps teams detect unexpected changes in timeseries-like data, whether it comes from metrics, logs-derived metrics, or traces-derived metrics. The goal is to simplify alerting, reduce maintenance overhead, and improve signal quality across the observability stack. [Learn more from the documentation](https://docs.victoriametrics.com/anomaly-detection/).

In general, Q1 2026 was less about a single foundational launch and more about improving the full `vmanomaly` workflow that was outlined in the [2025 recap](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2025/): simpler product onboarding, better UI ergonomics, AI assistance directly inside the UI, broader playground coverage, and more model-agnostic tools for reducing false positives in production.

### Q1 Highlights

The quarter continued the direction established in [2025](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2025/): 

**make anomaly detection easier to try, operate, and adapt to real-world use cases - from PoC to production**.

In Q1'26, `vmanomaly` shipped:

- 8 service releases, from [`v1.28.4`](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1284) to [`v1.29.3`](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1293)
- 7 [UI](https://docs.victoriametrics.com/anomaly-detection/ui/) iterations, from [`v1.4.1`](https://docs.victoriametrics.com/anomaly-detection/ui/#v141) to [`v1.6.1`](https://docs.victoriametrics.com/anomaly-detection/ui/#v161), including new features, easier authentication flows and bug fixes.
- [AI copilot integration inside the UI](#ai-assistance-inside-the-ui), allowing the assistant to see and interact with the current UI state for better suggestions and guidance
- A third public [playground](https://docs.victoriametrics.com/anomaly-detection/ui/#playgrounds) for **traces**, joining the existing metrics and logs playgrounds
- A new **model-agnostic** argument for reducing false positives via [relative deviation from expected value](https://docs.victoriametrics.com/anomaly-detection/components/models/#minimal-relative-deviation-from-expected)
- Documentation updates around the `server` [section](https://docs.victoriametrics.com/anomaly-detection/components/server/) of the config for fine-tuning UI appearance for internal users
- A broad batch of bug fixes and resource optimizations for models, writer, and the UI/server components

> [!TIP] Find video overview from our [Q1'26 meetup](https://youtu.be/Q76cd8gUtvY?t=4920) for more details.
> <iframe src="https://www.youtube.com/embed/Q76cd8gUtvY?start=4920" title="VictoriaMetrics Anomaly Detection Q1 2026 Updates" width="100%" height="400" frameborder="1" allowfullscreen></iframe>

### UI Improvements

The [UI](https://docs.victoriametrics.com/anomaly-detection/ui/) kept receiving steady improvements during Q1'26. The main theme here was making backtesting and configuration finetuning smoother for day-to-day users:

- Sorting and filtering of series returned by "Execute Query" and "Detect Anomalies" (click on "Table View" toggle in the legend or add `&legend_view=table` to the [default state](https://docs.victoriametrics.com/anomaly-detection/ui/#default-state) of the UI)
  
  ![filtering-and-sorting](/blog/victoriametrics-anomaly-detection-updates-2026-q1/vmanomaly-ui-filtering-sorting.webp)

- Anomaly streaks (consecutive anomalies) are now plotted as line segments instead of individual points, making it easier to visually identify longer-term patterns

    ![anomaly-streaks](/blog/victoriametrics-anomaly-detection-updates-2026-q1/vmanomaly-ui-streaks-line.png)

- Added support for UI to use reader-wise settings (credentials, TLS, etc.) for connecting to data sources, which otherwise would require having vmauth in front of both UI and data sources, just for the UI to be able to connect to them. See [Authentication](https://docs.victoriametrics.com/anomaly-detection/ui/#authentication) section for details. See config example below:

    ```yaml
    server:
      path_prefix: '/vmanomaly'  # optional path prefix for all HTTP routes
      
      # see https://docs.victoriametrics.com/anomaly-detection/ui/#default-state section for details on constructing the value from UI state
      ui_default_state: '#/?anomaly_threshold=1.0&anomaly_consecutive=true&fit_window=3d'  # optional default UI state opened on /vmui/
      max_concurrent_tasks: 4  # maximum number of concurrent anomaly detection tasks processed by backend

      use_reader_connection_settings: true  # if set to true, UI will use connection settings from reader configuration below when connecting to data sources, allowing it to connect with the same credentials, TLS settings, etc. without requiring having vmauth in front of both UI and data sources.

    # other vmanomaly configuration sections, like reader, scheduler, models, etc.
    reader:
      datasource_url: %{DS_URL}
      user: %{DS_USER}
      password: %{DS_PASSWORD}
      # or
      # bearer_token: %{DS_BEARER_TOKEN}
      verify_tls: false
      # other reader settings
    ```

- Better visualization of business arguments and how they affect anomaly detection behavior, such as detection direction (upward/downward blue area), expected value, and deviation thresholds (yellow bonds around the expected value). This is turned on by toggling the "Business Arguments" option in the legend or adding `&show_business_boundaries=true` to the [default state](https://docs.victoriametrics.com/anomaly-detection/ui/#default-state) of the UI.

  ![business-arguments](/blog/victoriametrics-anomaly-detection-updates-2026-q1/vmanomaly-ui-business-arguments.webp)

This extends the [UI work introduced in Q4'2025](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2025/#ui), but shifts the emphasis from "the UI exists" to "the UI is becoming a more practical daily workflow tool based on regular user feedback".

### AI Assistance Inside the UI

One of the most notable Q1'26 additions is **[AI copilot](https://docs.victoriametrics.com/anomaly-detection/ui/#ai-assistance) inside the UI**. Starting from [v1.29.3](https://docs.victoriametrics.com/anomaly-detection/changelog/#v1293), copilot can see and interact with the current UI state, making it possible to suggest **contextually better**

- models and their hyperparameters
- data ranges and step sizes for backtesting
- scheduling settings
- query configurations (using up-to-date MetricsQL and LogsQL syntax)

While doing so, it can also provide full documentation Q&A including UI elements and configuration options.


> [!TIP]
When connected to the [vmanomaly MCP server](https://github.com/VictoriaMetrics-Community/mcp-vmanomaly), the assistant can also use [the toolset](https://github.com/VictoriaMetrics/mcp-vmanomaly#toolset), such as config validation, documentation Q&A, and alerting-rule generation.

This is an important step in turning the MCP integration from a separate power-user tool into part of the main product workflow, **so that all users can benefit from AI assistance without needing to switch in-between interfaces**.

Here's a quick demo of how it takes server configuration and the current UI state into account when adjusting the state to requested changes:

<iframe
src="/blog/victoriametrics-anomaly-detection-updates-2026-q1/ai-copilot.webm"
width="100%"
height="640"
frameborder="0"
allowfullscreen
></iframe>

### Reducing False Positives Further

Another Q1'26 direction was further reducing false positives with simple business-aware controls. Service gained a [new model-agnostic argument](https://docs.victoriametrics.com/anomaly-detection/components/models/#minimal-relative-deviation-from-expected) for working with **relative deviation from expected value**, complementing the broader effort to make anomaly detection decisions closer to business significance rather than raw statistical distance alone. Works well for queries which **return metrics of unknown in advance scale, where even small absolute deviations can be significant**, and where confidence intervals produced by Machine Learning part, may not fully capture the expected variability of the metric.

<details>
<summary>Example Use Case (Unfold to see details)</summary>

**Example**: Consider a scenario of monitoring incoming traffic to websites, and the volume is usually *unknown in advance* (from tens to thousands visitors per second). 
- Setting *absolute* deviation threshold with `min_dev_from_expected` may not work well for reducing false positives, as it will be either too high to fully hide small in absolute number (yet big in relative change) deviations for low-traffic websites, or too low to avoid false positives for high-traffic websites.
- Instead, setting *relative thresholds* with `min_rel_dev_from_expected` tto values like (20, 40) - will ensure that traffic changes less than 20% drop or 40% increase won't be considered anomalous, even if they exceed confidence intervals, thus, reducing false positives for small absolute deviations that are relatively significant.

</details>

> [!TIP] Minimal relative and absolute deviation thresholds **can be combined together for better results, and can be set separately for upward and downward directions**. For example, setting `min_dev_from_expected` to (100, 0) and `min_rel_dev_from_expected` to (20, 0) will ensure that deviations should be both more than 100 in absolute number AND more than 20% in relative change to be considered anomalous in the downward direction, while for upward anomalies, there are no minimal deviation thresholds set, meaning that any increase above the upper confidence interval of Machine Learning model will be considered anomalous.

```yaml
# other components like writer, schedulers, monitoring ...
reader:
  # ...
  queries:
    # the usage of min_rel_dev should reduce false positives here
    need_to_include_min_rel_dev: 
      expr: small_abs_values_metricsql_expression
    # min_rel_dev is not really needed here
    normal_behavior: 
      expr: no_need_to_exclude_small_deviations_metricsql_expression
models:
  zscore_with_min_rel_dev:
    class: 'zscore_online'
    z_threshold: 3
    min_rel_dev_from_expected: [10, 20]  # set different thresholds for both directions, meaning that relative deviations less than 10% when y < yhat and less than 20% when y > yhat won't be considered anomalous, even if they exceed confidence intervals, thus, reducing false positives for small absolute deviations that are relatively significant
    queries: ['need_to_include_min_rel_dev']  # use such models on queries where domain experience confirm usefulness
  zscore_wo_min_rel_dev:
    class: 'zscore_online'
    z_threshold: 3
    # if not set, equals to setting min_rel_dev_from_expected == 0 (meaning no filtering is applied)
    # min_rel_dev_from_expected: [0, 0]
    queries: ['normal_behavior']  # use the default where it's not needed
```

### Public Playground for Traces

In [2025](https://victoriametrics.com/blog/victoriametrics-anomaly-detection-updates-2025/#public-playground), we launched public playgrounds for trying `vmanomaly` without a full enterprise deployment. Q1'26 extends that with the third playground, this time focused on traces-derived metrics.

Public playgrounds now cover:

- Metrics: https://play-vmanomaly.victoriametrics.com/metrics/vmui/
- Logs: https://play-vmanomaly.victoriametrics.com/logs/vmui/
- Traces: https://play-vmanomaly.victoriametrics.com/traces/vmui/

All three expose a [pre-filled default UI state](https://docs.victoriametrics.com/anomaly-detection/ui/#default-ui-state), which makes it easier to open the playground and immediately explore how anomaly detection behaves on example data.

> [!TIP] Visit new [playground overview page](https://docs.victoriametrics.com/playgrounds/) to see all the playgrounds for VictoriaMetrics ecosystem in one place.

### Documentation and Configuration Polish

Q1'26 also included documentation updates around the `server` [section](https://docs.victoriametrics.com/anomaly-detection/config/#server) of the configuration. Now it's easier to fine-tune the UI appearance and related behavior for internal users. This is especially helpful for teams that want to use `vmanomaly` as an internal tool for exploring anomalies without necessarily exposing every detail to end-users, by adjusting the default UI state, setting proper authentication flows, path prefixes, and more.

### Bug Fixes and Optimizations

Not every useful quarterly change deserves a flashy headline. To make it short, Q1'26 releases also included a broad batch of bug fixes and resource optimizations for models, writer throughput, and the UI/server components.

### Work in Progress

We have several work-in-progress items that are expected to be released in the next few weeks, including:

- Agentic skills for `vmanomaly`, to become part of the [public skills repository](https://github.com/VictoriaMetrics/skills)

  > [!TIP] See actual list of AI tools for VictoriaMetrics ecosystem [here](https://docs.victoriametrics.com/ai-tools/)

- On our path to make all the models [online](https://docs.victoriametrics.com/anomaly-detection/components/models/#online-models) - able to learn from streaming data without the need for repetitive retraining - an online Prophet-like model combining flexibility with better resource efficiency is under development and expected to be released in the next few weeks to replace [offline Prophet](https://docs.victoriametrics.com/anomaly-detection/components/models/#prophet) in the configuration. This will allow to **have a single model for a wider range of use cases**, without the need to choose between the flexibility of Prophet and the efficiency of more basic online models.

### What's Next?

The roadmap direction for Q2-Q4'26, includes:

- Updates to the `vmanomaly` MCP server and skills (once released) for easier product onboarding and cross-product interaction
- Dynamic presets with metric-subset discovery for sources such as Kubernetes, VictoriaMetrics, and `node-exporter`, allowing to set up anomaly detection with a few clicks on the most important signals for common use cases (configs, alerting rules, etc.)
- Root-cause analysis for assisted causal drill-downs using logs, traces, and metrics together
- Integration with [VictoriaMetrics Cloud](https://victoriametrics.com/products/cloud/)

Stay tuned!

### Frequently Asked Questions

**What were the main VictoriaMetrics Anomaly Detection updates in Q1 2026?** In Q1 2026, VictoriaMetrics Anomaly Detection focused on making vmanomaly easier to try, tune, operate, and move into production. Key updates included 8 service releases, 7 anomaly detection UI iterations, AI copilot integration inside the UI, a new public playground for traces, new controls for reducing false positives, documentation improvements, and resource optimizations across models, writer, UI, and server components.

**What is vmanomaly?** vmanomaly is the VictoriaMetrics Anomaly Detection service. It is designed to scan time-series data from VictoriaMetrics, VictoriaLogs, or VictoriaTraces and generate anomaly scores based on expected behavior. These anomaly scores can then be stored back in VictoriaMetrics and used in dashboards, alerts, and downstream observability workflows. The goal is to reduce manual threshold tuning while improving the reliability of anomaly detection.

**What is anomaly detection machine learning?** Anomaly detection machine learning uses models to learn expected behavior from time-series data, then identify data points that deviate from that expected pattern. In VictoriaMetrics Anomaly Detection, vmanomaly continuously applies machine learning models to metrics, logs-derived metrics, or traces-derived metrics and produces anomaly scores that can be used for alerting, visualization, and analysis. This helps teams move beyond static thresholds and detect abnormal system behavior more reliably.

**How does anomaly detection improve data quality?** Anomaly detection improves data quality by making unusual, unexpected, or potentially incorrect patterns easier to identify. In observability workflows, this does not necessarily mean the raw data is changed or "cleaned". Instead, VictoriaMetrics Anomaly Detection helps improve signal quality by detecting deviations from expected behavior, reducing alert noise, and making it easier for teams to distinguish meaningful anomalies from normal variation, seasonality, or trend changes.

**How does the anomaly detection UI help teams tune models before production?** The anomaly detection UI helps teams explore input data, configure models, visualize predictions and anomaly scores, and test how model settings affect results before deploying those configurations in production. In Q1 2026, VictoriaMetrics continued improving the UI with better query result handling, business-argument visualization, AI assistance, and tools for tuning false-positive behavior. This makes it easier for teams to validate models, apply domain knowledge, and generate production-ready configurations with more confidence.

---

*I am Fred Navruzov and have previously worked as a Data Science Lead at Ukrainian and international companies such as Competera, Rozetka, and TiltingPoint. Since 2023, I have been leading the Anomaly Detection team at VictoriaMetrics.*
