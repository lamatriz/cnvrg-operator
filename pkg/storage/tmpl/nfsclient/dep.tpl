kind: Deployment
apiVersion: apps/v1
metadata:
  name: nfs-client-provisioner
  namespace: {{ ns . }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nfs-client-provisioner
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: nfs-client-provisioner
    spec:
      {{- if isTrue .Spec.Tenancy.Enabled }}
      nodeSelector:
        "{{ .Spec.Tenancy.Key }}": "{{ .Spec.Tenancy.Value }}"
      tolerations:
        - key: "{{ .Spec.Tenancy.Key }}"
          operator: "Equal"
          value: "{{ .Spec.Tenancy.Value }}"
          effect: "NoSchedule"
      {{- end }}
      serviceAccountName: nfs-client-provisioner
      containers:
        - name: nfs-client-provisioner
          image: {{ .Spec.Storage.Nfs.Image }}
          volumeMounts:
            - name: nfs-client-root
              mountPath: /persistentvolumes
          env:
            - name: PROVISIONER_NAME
              value: {{ .Spec.Storage.Nfs.Provisioner }}
            - name: NFS_SERVER
              value: {{ .Spec.Storage.Nfs.Server }}
            - name: NFS_PATH
              value: "{{ .Spec.Storage.Nfs.Path }}"
          resources:
            limits:
              cpu: {{ .Spec.Storage.Nfs.CPULimit }}
              memory: {{ .Spec.Storage.Nfs.MemoryLimit }}
            requests:
              cpu: {{ .Spec.Storage.Nfs.CPURequest }}
              memory: {{ .Spec.Storage.Nfs.MemoryRequest }}
      volumes:
        - name: nfs-client-root
          nfs:
            server: {{ .Spec.Storage.Nfs.Server }}
            path: {{ .Spec.Storage.Nfs.Path }}