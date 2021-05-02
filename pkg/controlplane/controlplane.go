package controlplane

import (
	mlopsv1 "github.com/cnvrg-operator/api/v1"
	"github.com/cnvrg-operator/pkg/desired"
	"github.com/markbates/pkger"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
)

const path = "/pkg/controlplane/tmpl"

func rbacState() []*desired.State {
	return []*desired.State{
		{

			TemplatePath:   path + "/conf/rbac/role.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.RoleGVR],
			Own:            true,
		},
		{

			TemplatePath:   path + "/conf/rbac/rolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.RoleBindingGVR],
			Own:            true,
		},
		{

			TemplatePath:   path + "/conf/rbac/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
	}
}

func controlPlaneConfigState() []*desired.State {
	return []*desired.State{
		{

			TemplatePath:   path + "/conf/cm/config-base.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ConfigMapGVR],
			Own:            true,
		},
		{

			TemplatePath:   path + "/conf/cm/config-networking.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ConfigMapGVR],
			Own:            true,
		},
		{

			TemplatePath:   path + "/conf/cm/secret-base.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
		{

			TemplatePath:   path + "/conf/cm/secret-ldap.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/conf/cm/secret-object-storage.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/conf/cm/secret-smtp.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
	}
}

func webAppState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/webapp/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/webapp/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/webapp/oauth.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/webapp/pdb.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PodDisruptionBudgetGVR],
			Own:            true,
		},
	}
}

func webAppIstioVs() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/webapp/vs.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IstioVsGVR],
			Own:            true,
		},
	}
}

func webAppOcpRoute() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/webapp/route.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.OcpRouteGVR],
			Own:            true,
		},
	}
}

func webAppIngress() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/webapp/ingress.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IngressGVR],
			Own:            true,
		},
	}
}

func sidekiqState() []*desired.State {
	return []*desired.State{
		{

			TemplatePath:   path + "/sidekiqs/sidekiq.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/sidekiqs/sidekiq-pdb.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PodDisruptionBudgetGVR],
			Own:            true,
		},
	}
}

func searchkiqState() []*desired.State {
	return []*desired.State{
		{

			TemplatePath:   path + "/sidekiqs/searchkiq.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/sidekiqs/searchkiq-pdb.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PodDisruptionBudgetGVR],
			Own:            true,
		},
	}
}

func systemkiqState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/sidekiqs/systemkiq.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/sidekiqs/systemkiq-pdb.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.PodDisruptionBudgetGVR],
			Own:            true,
		},
	}
}

func hyperState() []*desired.State {
	return []*desired.State{
		{

			TemplatePath:   path + "/hyper/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
		{

			TemplatePath:   path + "/hyper/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            true,
		},
	}
}

func ssoState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/webapp/oauth.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
	}
}

func mpiAppState() []*desired.State {
	return []*desired.State{

		{
			TemplatePath:   path + "/mpi/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/mpi/rolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.RoleBindingGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/mpi/secret.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SecretGVR],
			Own:            true,
		},
		{
			TemplatePath:   path + "/mpi/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            true,
		},
	}
}

func mpiInfraState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/mpi/clusterrole.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleGVR],
			Own:            true,
		},
	}
}

func State(cnvrgApp *mlopsv1.CnvrgApp) []*desired.State {
	var state []*desired.State
	state = append(state, rbacState()...)
	state = append(state, controlPlaneConfigState()...)

	if *cnvrgApp.Spec.ControlPlane.WebApp.Enabled {
		state = append(state, webAppState()...)
		if cnvrgApp.Spec.Networking.Ingress.Type == mlopsv1.IstioIngress {
			state = append(state, webAppIstioVs()...)
		}
		if cnvrgApp.Spec.Networking.Ingress.Type == mlopsv1.OpenShiftIngress {
			state = append(state, webAppOcpRoute()...)
		}

		if cnvrgApp.Spec.Networking.Ingress.Type == mlopsv1.NginxIngress {
			state = append(state, webAppIngress()...)
		}
	}

	if *cnvrgApp.Spec.SSO.Enabled {
		state = append(state, ssoState()...)
	}

	if *cnvrgApp.Spec.ControlPlane.Sidekiq.Enabled && *cnvrgApp.Spec.ControlPlane.Sidekiq.Split {
		state = append(state, sidekiqState()...)
	}

	if *cnvrgApp.Spec.ControlPlane.Searchkiq.Enabled && *cnvrgApp.Spec.ControlPlane.Sidekiq.Split {
		state = append(state, searchkiqState()...)
	}

	if *cnvrgApp.Spec.ControlPlane.Systemkiq.Enabled && *cnvrgApp.Spec.ControlPlane.Sidekiq.Split {
		state = append(state, systemkiqState()...)
	}

	if *cnvrgApp.Spec.ControlPlane.Sidekiq.Enabled && *cnvrgApp.Spec.ControlPlane.Sidekiq.Split {
		state = append(state, sidekiqState()...)
	}

	if *cnvrgApp.Spec.ControlPlane.Hyper.Enabled {
		state = append(state, hyperState()...)
	}

	if *cnvrgApp.Spec.ControlPlane.Mpi.Enabled {
		state = append(state, mpiAppState()...)
	}
	return state
}

func MpiInfraState() []*desired.State {
	return mpiInfraState()
}

func Crds() (crds []*desired.State) {
	err := pkger.Walk(path+"/mpi/crds", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		crd := &desired.State{

			TemplatePath:   path,
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.CrdGVR],
			Own:            false,
		}
		crds = append(crds, crd)
		return nil
	})
	if err != nil {
		zap.S().Error(err, "error loading mpi crds")
	}
	return
}
