---
title : "Terraform"
menu:
  docs:
    parent: "integrations"
---

[Terraform](https://www.terraform.io/) is an infrastructure-as-code tool that allows you to define and
provision infrastructure using declarative configuration files. The VictoriaMetrics Cloud Terraform provider
enables you to manage your VictoriaMetrics Cloud infrastructure programmatically, including deployments,
access tokens, and alerting rules.

This integration allows you to version control your monitoring infrastructure, automate deployments, and
ensure consistent configuration across environments using standard Terraform workflows.

## Integrating with Terraform

All VictoriaMetrics Cloud Terraform provider operations require an API key for authentication. The
configuration examples below contain a placeholder `<YOUR_API_KEY>` that needs to be replaced with your
actual API key.

More information about API keys you can find in the
[API documentation](https://docs.victoriametrics.com/victoriametrics-cloud/api/).

The Terraform provider is available in the
[Terraform Registry](https://registry.terraform.io/providers/VictoriaMetrics/victoriametricscloud/latest/docs),
and will be automatically downloaded when you run `terraform init`.

To set up the Terraform provider for VictoriaMetrics Cloud, visit the
[Terraform integration](https://console.victoriametrics.cloud/integrations/terraform) in the Cloud Console
or follow this interactive guide:

<iframe
    width="100%"
    height="2800"
    name="iframe"
    id="integration"
    frameborder="0"
    src="https://console.victoriametrics.cloud/public/integrations/terraform"
    style="background: white;" >
</iframe>
