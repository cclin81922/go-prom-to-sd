# Usage

```
git clone https://github.com/cclin81922/go-prom-to-sd.git
cd go-prom-to-sd

bash BuildingwithCloudNativeBuildpacks.sh

yq w -i daemonset.yaml "spec.containers[0].image" "gcr.io/$(gcloud config get-value project)/prom-to-sd"
kubectl apply -f daemonset.yaml
```

# About Instrumenting a GO Application for Prometheus

https://prometheus.io/docs/guides/go-application/

# About prometheus-to-sd

https://github.com/GoogleCloudPlatform/k8s-stackdriver/tree/master/prometheus-to-sd

# About yq

https://learnk8s.io/templating-yaml-with-code
