apiVersion: apps/v1
kind: Deployment
metadata:
  name: searchkiq
  namespace: {{ .CnvrgNs }}
  labels:
    app: searchkiq
spec:
  replicas: {{ .ControlPlan.Searchkiq.Replicas }}
  selector:
    matchLabels:
      app: searchkiq
  template:
    metadata:
      labels:
        app: searchkiq
    spec:
      {{- if eq .ControlPlan.Tenancy.Enabled "true" }}
      nodeSelector:
        {{ .ControlPlan.Tenancy.Key }}: "{{ .ControlPlan.Tenancy.Value }}"
      {{- end }}
      tolerations:
        - key: "{{ .ControlPlan.Tenancy.Key }}"
          operator: "Equal"
          value: "{{ .ControlPlan.Tenancy.Value }}"
          effect: "NoSchedule"
      serviceAccountName: {{ .ControlPlan.Rbac.ServiceAccountName }}
      terminationGracePeriodSeconds: {{ .ControlPlan.Searchkiq.KillTimeout }}
      containers:
      - name: sidekiq
        image: {{ .ControlPlan.WebApp.Image}}
        env:
        - name: "CNVRG_RUN_MODE"
          value: "sidekiq"
        - name: "SIDEKIQ_SEARCHKICK"
          value: "true"
        imagePullPolicy: Always
        {{- if eq .ControlPlan.ObjectStorage.CnvrgStorageType "gcp" }}
        volumeMounts:
          - name: "{{ .ControlPlan.ObjectStorage.GcpStorageSecret }}"
            mountPath: "{{ .ControlPlan.ObjectStorage.GcpKeyfileMountPath }}"
            readOnly: true
        {{- end }}
        envFrom:
          - configMapRef:
              name: cp-base-config
          - configMapRef:
              name: cp-networking-config
          - secretRef:
              name: cp-base-secret
          - secretRef:
              name: cp-ldap
          - secretRef:
              name: cp-object-storage
          - secretRef:
              name: {{ .Pg.SvcName }}
        resources:
          requests:
            cpu: {{ .ControlPlan.Searchkiq.CPU }}
            memory: {{ .ControlPlan.Searchkiq.Memory }}
        lifecycle:
          preStop:
            exec:
              command: ["/bin/bash","-lc","sidekiqctl quiet sidekiq-0.pid && sidekiqctl stop sidekiq-0.pid {{ .ControlPlan.Searchkiq.KillTimeout }}"]
      {{- if eq .ControlPlan.ObjectStorage.CnvrgStorageType "gcp" }}
      volumes:
        - name: {{ .ControlPlan.ObjectStorage.GcpStorageSecret }}
          secret:
            secretName: {{ .ControlPlan.ObjectStorage.GcpStorageSecret }}
      {{- end }}
      initContainers:
      - name: seeder
        image: {{.ControlPlan.Seeder.Image}}
        command: ["/bin/bash", "-c", "python3 cnvrg-boot.py seeder --mode worker"]
        env:
        - name: "CNVRG_NS"
          value: {{ .CnvrgNs }}