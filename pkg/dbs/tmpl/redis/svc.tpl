apiVersion: v1
kind: Service
metadata:
  name: {{.Spec.Dbs.Redis.SvcName}}
  namespace: {{ ns . }}
  labels:
    app: {{.Spec.Dbs.Redis.SvcName }}
    owner: cnvrg-control-plane
spec:
  ports:
  - name: redis
    port: {{ .Spec.Dbs.Redis.Port }}
  selector:
    app: {{ .Spec.Dbs.Redis.SvcName }}