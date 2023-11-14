[![Twitter](https://img.shields.io/twitter/follow/appscodehq.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=AppsCodeHQ)

# Kluster API installer

Kluster API Helm charts

## Instructions

```bash
# in hub cluster
helm install cluster-manager-hub charts/cluster-manager-hub \
  -n open-cluster-management --create-namespace

# in spoke cluster
helm install cluster-manager-spoke charts/cluster-manager-spoke \
  -n open-cluster-management --create-namespace \
  --set clusterName=*** \
  --set hub.apiServer=https://***.com \
  --set hub.token=***
```

## Support

To speak with us, please leave a message on [our website](https://appscode.com/contact/).

To receive product annoucements, follow us on [Twitter](https://twitter.com/AppsCodeHQ).
