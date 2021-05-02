apiVersion: v1
kind: ServiceAccount
metadata:
  name: cnvrg-infra-prometheus
  namespace: {{ ns . }}
  labels:
    owner: cnvrg-control-plane
imagePullSecrets:
  - name: {{ .Spec.Registry.Name }}
