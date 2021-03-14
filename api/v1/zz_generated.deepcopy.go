// +build !ignore_autogenerated

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseConfig) DeepCopyInto(out *BaseConfig) {
	*out = *in
	if in.FeatureFlags != nil {
		in, out := &in.FeatureFlags, &out.FeatureFlags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseConfig.
func (in *BaseConfig) DeepCopy() *BaseConfig {
	if in == nil {
		return nil
	}
	out := new(BaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cnvrg) DeepCopyInto(out *Cnvrg) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cnvrg.
func (in *Cnvrg) DeepCopy() *Cnvrg {
	if in == nil {
		return nil
	}
	out := new(Cnvrg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgApp) DeepCopyInto(out *CnvrgApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgApp.
func (in *CnvrgApp) DeepCopy() *CnvrgApp {
	if in == nil {
		return nil
	}
	out := new(CnvrgApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnvrgApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppList) DeepCopyInto(out *CnvrgAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CnvrgApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppList.
func (in *CnvrgAppList) DeepCopy() *CnvrgAppList {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CnvrgAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppSpec) DeepCopyInto(out *CnvrgAppSpec) {
	*out = *in
	in.ControlPlan.DeepCopyInto(&out.ControlPlan)
	out.Pg = in.Pg
	out.Storage = in.Storage
	out.Networking = in.Networking
	out.Minio = in.Minio
	out.Redis = in.Redis
	out.Logging = in.Logging
	out.Monitoring = in.Monitoring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppSpec.
func (in *CnvrgAppSpec) DeepCopy() *CnvrgAppSpec {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgAppStatus) DeepCopyInto(out *CnvrgAppStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgAppStatus.
func (in *CnvrgAppStatus) DeepCopy() *CnvrgAppStatus {
	if in == nil {
		return nil
	}
	out := new(CnvrgAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CnvrgRouter) DeepCopyInto(out *CnvrgRouter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CnvrgRouter.
func (in *CnvrgRouter) DeepCopy() *CnvrgRouter {
	if in == nil {
		return nil
	}
	out := new(CnvrgRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsistentHash) DeepCopyInto(out *ConsistentHash) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsistentHash.
func (in *ConsistentHash) DeepCopy() *ConsistentHash {
	if in == nil {
		return nil
	}
	out := new(ConsistentHash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlan) DeepCopyInto(out *ControlPlan) {
	*out = *in
	out.WebApp = in.WebApp
	out.Sidekiq = in.Sidekiq
	out.Searchkiq = in.Searchkiq
	out.Systemkiq = in.Systemkiq
	out.CnvrgRouter = in.CnvrgRouter
	out.Hyper = in.Hyper
	out.Seeder = in.Seeder
	in.BaseConfig.DeepCopyInto(&out.BaseConfig)
	out.Ldap = in.Ldap
	out.Registry = in.Registry
	out.Rbac = in.Rbac
	out.SMTP = in.SMTP
	out.Tenancy = in.Tenancy
	in.OauthProxy.DeepCopyInto(&out.OauthProxy)
	out.ObjectStorage = in.ObjectStorage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlan.
func (in *ControlPlan) DeepCopy() *ControlPlan {
	if in == nil {
		return nil
	}
	out := new(ControlPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DcgmExporter) DeepCopyInto(out *DcgmExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DcgmExporter.
func (in *DcgmExporter) DeepCopy() *DcgmExporter {
	if in == nil {
		return nil
	}
	out := new(DcgmExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultServiceMonitors) DeepCopyInto(out *DefaultServiceMonitors) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultServiceMonitors.
func (in *DefaultServiceMonitors) DeepCopy() *DefaultServiceMonitors {
	if in == nil {
		return nil
	}
	out := new(DefaultServiceMonitors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elastalert) DeepCopyInto(out *Elastalert) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elastalert.
func (in *Elastalert) DeepCopy() *Elastalert {
	if in == nil {
		return nil
	}
	out := new(Elastalert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Es) DeepCopyInto(out *Es) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Es.
func (in *Es) DeepCopy() *Es {
	if in == nil {
		return nil
	}
	out := new(Es)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fluentd) DeepCopyInto(out *Fluentd) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fluentd.
func (in *Fluentd) DeepCopy() *Fluentd {
	if in == nil {
		return nil
	}
	out := new(Fluentd)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grafana) DeepCopyInto(out *Grafana) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grafana.
func (in *Grafana) DeepCopy() *Grafana {
	if in == nil {
		return nil
	}
	out := new(Grafana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPS) DeepCopyInto(out *HTTPS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPS.
func (in *HTTPS) DeepCopy() *HTTPS {
	if in == nil {
		return nil
	}
	out := new(HTTPS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hostpath) DeepCopyInto(out *Hostpath) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hostpath.
func (in *Hostpath) DeepCopy() *Hostpath {
	if in == nil {
		return nil
	}
	out := new(Hostpath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HugePages) DeepCopyInto(out *HugePages) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HugePages.
func (in *HugePages) DeepCopy() *HugePages {
	if in == nil {
		return nil
	}
	out := new(HugePages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hyper) DeepCopyInto(out *Hyper) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hyper.
func (in *Hyper) DeepCopy() *Hyper {
	if in == nil {
		return nil
	}
	out := new(Hyper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdleMetricsExporter) DeepCopyInto(out *IdleMetricsExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdleMetricsExporter.
func (in *IdleMetricsExporter) DeepCopy() *IdleMetricsExporter {
	if in == nil {
		return nil
	}
	out := new(IdleMetricsExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Images) DeepCopyInto(out *Images) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Images.
func (in *Images) DeepCopy() *Images {
	if in == nil {
		return nil
	}
	out := new(Images)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Istio) DeepCopyInto(out *Istio) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Istio.
func (in *Istio) DeepCopy() *Istio {
	if in == nil {
		return nil
	}
	out := new(Istio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kibana) DeepCopyInto(out *Kibana) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kibana.
func (in *Kibana) DeepCopy() *Kibana {
	if in == nil {
		return nil
	}
	out := new(Kibana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStateMetrics) DeepCopyInto(out *KubeStateMetrics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStateMetrics.
func (in *KubeStateMetrics) DeepCopy() *KubeStateMetrics {
	if in == nil {
		return nil
	}
	out := new(KubeStateMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ldap) DeepCopyInto(out *Ldap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ldap.
func (in *Ldap) DeepCopy() *Ldap {
	if in == nil {
		return nil
	}
	out := new(Ldap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Limits) DeepCopyInto(out *Limits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Limits.
func (in *Limits) DeepCopy() *Limits {
	if in == nil {
		return nil
	}
	out := new(Limits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	out.Es = in.Es
	out.Elastalert = in.Elastalert
	out.Fluentd = in.Fluentd
	out.Kibana = in.Kibana
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsServer) DeepCopyInto(out *MetricsServer) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsServer.
func (in *MetricsServer) DeepCopy() *MetricsServer {
	if in == nil {
		return nil
	}
	out := new(MetricsServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Minio) DeepCopyInto(out *Minio) {
	*out = *in
	out.SharedStorage = in.SharedStorage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Minio.
func (in *Minio) DeepCopy() *Minio {
	if in == nil {
		return nil
	}
	out := new(Minio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinioExporter) DeepCopyInto(out *MinioExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinioExporter.
func (in *MinioExporter) DeepCopy() *MinioExporter {
	if in == nil {
		return nil
	}
	out := new(MinioExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	out.PrometheusOperator = in.PrometheusOperator
	out.Prometheus = in.Prometheus
	out.NodeExporter = in.NodeExporter
	out.KubeStateMetrics = in.KubeStateMetrics
	out.Grafana = in.Grafana
	out.DefaultServiceMonitors = in.DefaultServiceMonitors
	out.SidekiqExporter = in.SidekiqExporter
	out.MinioExporter = in.MinioExporter
	out.DcgmExporter = in.DcgmExporter
	out.IdleMetricsExporter = in.IdleMetricsExporter
	out.MetricsServer = in.MetricsServer
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networking) DeepCopyInto(out *Networking) {
	*out = *in
	out.HTTPS = in.HTTPS
	out.Istio = in.Istio
	out.Ingress = in.Ingress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networking.
func (in *Networking) DeepCopy() *Networking {
	if in == nil {
		return nil
	}
	out := new(Networking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nfs) DeepCopyInto(out *Nfs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nfs.
func (in *Nfs) DeepCopy() *Nfs {
	if in == nil {
		return nil
	}
	out := new(Nfs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeExporter) DeepCopyInto(out *NodeExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeExporter.
func (in *NodeExporter) DeepCopy() *NodeExporter {
	if in == nil {
		return nil
	}
	out := new(NodeExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OauthProxy) DeepCopyInto(out *OauthProxy) {
	*out = *in
	if in.SkipAuthRegex != nil {
		in, out := &in.SkipAuthRegex, &out.SkipAuthRegex
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OauthProxy.
func (in *OauthProxy) DeepCopy() *OauthProxy {
	if in == nil {
		return nil
	}
	out := new(OauthProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorage) DeepCopyInto(out *ObjectStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorage.
func (in *ObjectStorage) DeepCopy() *ObjectStorage {
	if in == nil {
		return nil
	}
	out := new(ObjectStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pg) DeepCopyInto(out *Pg) {
	*out = *in
	out.HugePages = in.HugePages
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pg.
func (in *Pg) DeepCopy() *Pg {
	if in == nil {
		return nil
	}
	out := new(Pg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusOperator) DeepCopyInto(out *PrometheusOperator) {
	*out = *in
	out.Images = in.Images
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusOperator.
func (in *PrometheusOperator) DeepCopy() *PrometheusOperator {
	if in == nil {
		return nil
	}
	out := new(PrometheusOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rbac) DeepCopyInto(out *Rbac) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rbac.
func (in *Rbac) DeepCopy() *Rbac {
	if in == nil {
		return nil
	}
	out := new(Rbac)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.Limits = in.Limits
	out.Requests = in.Requests
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Requests) DeepCopyInto(out *Requests) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Requests.
func (in *Requests) DeepCopy() *Requests {
	if in == nil {
		return nil
	}
	out := new(Requests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMTP) DeepCopyInto(out *SMTP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMTP.
func (in *SMTP) DeepCopy() *SMTP {
	if in == nil {
		return nil
	}
	out := new(SMTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Searchkiq) DeepCopyInto(out *Searchkiq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Searchkiq.
func (in *Searchkiq) DeepCopy() *Searchkiq {
	if in == nil {
		return nil
	}
	out := new(Searchkiq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Seeder) DeepCopyInto(out *Seeder) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Seeder.
func (in *Seeder) DeepCopy() *Seeder {
	if in == nil {
		return nil
	}
	out := new(Seeder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedStorage) DeepCopyInto(out *SharedStorage) {
	*out = *in
	out.ConsistentHash = in.ConsistentHash
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedStorage.
func (in *SharedStorage) DeepCopy() *SharedStorage {
	if in == nil {
		return nil
	}
	out := new(SharedStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidekiq) DeepCopyInto(out *Sidekiq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidekiq.
func (in *Sidekiq) DeepCopy() *Sidekiq {
	if in == nil {
		return nil
	}
	out := new(Sidekiq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SidekiqExporter) DeepCopyInto(out *SidekiqExporter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SidekiqExporter.
func (in *SidekiqExporter) DeepCopy() *SidekiqExporter {
	if in == nil {
		return nil
	}
	out := new(SidekiqExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	out.Hostpath = in.Hostpath
	out.Nfs = in.Nfs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Systemkiq) DeepCopyInto(out *Systemkiq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Systemkiq.
func (in *Systemkiq) DeepCopy() *Systemkiq {
	if in == nil {
		return nil
	}
	out := new(Systemkiq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tenancy) DeepCopyInto(out *Tenancy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tenancy.
func (in *Tenancy) DeepCopy() *Tenancy {
	if in == nil {
		return nil
	}
	out := new(Tenancy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebApp) DeepCopyInto(out *WebApp) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebApp.
func (in *WebApp) DeepCopy() *WebApp {
	if in == nil {
		return nil
	}
	out := new(WebApp)
	in.DeepCopyInto(out)
	return out
}
