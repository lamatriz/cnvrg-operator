apiVersion: v1
kind: Secret
metadata:
  name: oauth-proxy-grafana
  namespace: {{ ns . }}
  annotations:
    {{- range $k, $v := .Spec.Annotations }}
    {{$k}}: "{{$v}}"
    {{- end }}
  labels:
    {{- range $k, $v := .Spec.Labels }}
    {{$k}}: "{{$v}}"
    {{- end }}
data:
  conf: {{ oauthProxyConfig . .Spec.Monitoring.Grafana.SvcName .Spec.Monitoring.Grafana.OauthProxy.SkipAuthRegex .Spec.SSO.Provider .Spec.Monitoring.Grafana.Port 3000 | b64enc }}