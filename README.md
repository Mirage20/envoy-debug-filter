# Envoy Debug Filter for Istio


### Apply the debug filter

    kubectl apply -f https://raw.githubusercontent.com/Mirage20/envoy-debug-filter/master/envoy-debug-filter.yaml

### Tail the debug logs to see what happen to your request
    kubectl logs $(kubectl get pod -l app=envoy-debug-filter -o jsonpath="{.items[0].metadata.name}") -f

### Remove the debug filter

    kubectl delete -f https://raw.githubusercontent.com/Mirage20/envoy-debug-filter/master/envoy-debug-filter.yaml
