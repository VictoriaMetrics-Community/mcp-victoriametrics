---
title : "CURL"
menu:
  docs:
    parent: "integrations"
---

You can use [curl](https://curl.se/) to interact with VictoriaMetrics Cloud for both **pushing** data and
**querying** stored data using HTTP API endpoints. This makes it a simple and flexible option for
testing or basic integrations.

For VictoriaTraces deployments, spans are pushed in OTLP/JSON format:

```sh
NOW=$(date +%s)
START="${NOW}000000000"
END="$((NOW + 1))000000000"

curl \
  -H 'Authorization: Bearer <YOUR_ACCESS_TOKEN>' \
  -H 'Content-Type: application/json' \
  -X POST '<DEPLOYMENT_ENDPOINT_URL>/insert/opentelemetry/v1/traces' \
  -d '{"resourceSpans":[{"resource":{"attributes":[{"key":"service.name","value":{"stringValue":"my-service"}}]},"scopeSpans":[{"scope":{},"spans":[{"traceId":"5b8aa5a2d2c872e8321cf37308d69df2","spanId":"051581bf3cb55c13","name":"curl-test-span","kind":1,"startTimeUnixNano":"'"$START"'","endTimeUnixNano":"'"$END"'"}]}]}]}'
```

> [!TIP]
> Spans carry their own timestamps, and VictoriaTraces uses the span end time as its
> [time field](https://docs.victoriametrics.com/victoriatraces/keyconcepts/#time-field). A span sent
> with a hardcoded timestamp from the past is stored at that past time, so it does not show up when
> looking at recent data, and it is
> [rejected altogether](https://docs.victoriametrics.com/victoriatraces/#retention) if it falls
> outside the deployment retention. That is why the snippet above builds the timestamps from `date`,
> producing a one second span that ends now and is immediately visible in the
> [Explore](https://docs.victoriametrics.com/victoriametrics-cloud/exploring-data/exploring-victoriatraces/)
> page.

Spans are read back through the Tempo API at `<DEPLOYMENT_ENDPOINT_URL>/select/tempo` or the Jaeger
API at `<DEPLOYMENT_ENDPOINT_URL>/select/jaeger`. For example, to list the traces of the service used
above:

```sh
curl -G \
  -H 'Authorization: Bearer <YOUR_ACCESS_TOKEN>' \
  '<DEPLOYMENT_ENDPOINT_URL>/select/jaeger/api/traces' \
  -d 'service=my-service'
```

## Integrating with CURL

All VictoriaMetrics Cloud integrations, including this one, require an access token for authentication.
The configuration examples below contain two placeholders: `<DEPLOYMENT_ENDPOINT_URL>` and
`<YOUR_ACCESS_TOKEN>`. These need to be replaced with your actual access token.

To generate your access token (with **write access** for pushing data or **read access** for querying),
follow the steps in the [Access Tokens documentation](https://docs.victoriametrics.com/victoriametrics-cloud/deployments/access-tokens/).

To integrate CURL with VictoriaMetrics Cloud, visit the [cloud console](https://console.victoriametrics.cloud/integrations/curl),
or simply follow this interactive guide:


{{< cloud-console-iframe "curl" >}}
