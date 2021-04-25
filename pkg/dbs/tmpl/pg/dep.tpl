apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Spec.Dbs.Pg.SvcName }}
  namespace: {{ ns . }}
  labels:
    app: {{.Spec.Dbs.Pg.SvcName }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.Spec.Dbs.Pg.SvcName }}
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: {{.Spec.Dbs.Pg.SvcName }}
    spec:
      serviceAccountName: {{ .Spec.Dbs.Pg.ServiceAccount }}
      securityContext:
        runAsUser: 26
        fsGroup: 26
      containers:
        - name: postgresql
          envFrom:
            - secretRef:
                name: {{ .Spec.Dbs.Pg.CredsRef }}
          image: {{.Spec.Dbs.Pg.Image}}
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: {{.Spec.Dbs.Pg.Port}}
              protocol: TCP
          livenessProbe:
            exec:
              command:
                - /usr/libexec/check-container
                - --live
            initialDelaySeconds: 120
            timeoutSeconds: 10
          readinessProbe:
            exec:
              command:
                - /usr/libexec/check-container
            initialDelaySeconds: 5
            timeoutSeconds: 1
          securityContext:
            capabilities: {}
            privileged: false
          terminationMessagePath: /dev/termination-log
          volumeMounts:
            - mountPath: /var/lib/pgsql/data
              name: postgres-data
            - mountPath: /dev/shm
              name: dshm
            {{- if eq .Spec.Dbs.Pg.HugePages.Enabled "true" }}
            - mountPath: "/hugepages"
              name: "hugepage"
            {{- end }}
          resources:
            {{- if eq .Spec.Dbs.Pg.HugePages.Enabled "true" }}
            limits:
              {{- if eq .Spec.Dbs.Pg.HugePages.Memory ""}}
              hugepages-{{ .Spec.Dbs.Pg.HugePages.Size }}: {{ .Spec.Dbs.Pg.MemoryRequest }}
              {{- else }}
              hugepages-{{ .Spec.Dbs.Pg.HugePages.Size }}: {{ .Spec.Dbs.Pg.HugePages.Memory }}
              {{- end }}
            {{- end }}
            requests:
              cpu: {{ .Spec.Dbs.Pg.CPURequest }}
              memory: {{ .Spec.Dbs.Pg.MemoryRequest }}
      volumes:
        - name: postgres-data
          persistentVolumeClaim:
            claimName: {{.Spec.Dbs.Pg.SvcName}}
        - name: dshm
          emptyDir:
            medium: Memory
            sizeLimit: 2Gi
        {{- if eq .Spec.Dbs.Pg.HugePages.Enabled "true" }}
        - name: "hugepage"
          emptyDir:
            medium: HugePages
        {{- end}}
