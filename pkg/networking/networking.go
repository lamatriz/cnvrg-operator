package networking

import (
	mlopsv1 "github.com/cnvrg-operator/api/v1"
	"github.com/cnvrg-operator/pkg/desired"
	"github.com/markbates/pkger"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
)

const path = "/pkg/networking/tmpl"

func istioInstanceState() []*desired.State {
	return []*desired.State{
		{

			TemplatePath:   path + "/istio/instance/clusterrole.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleGVR],
			Own:            false,
			Updatable:      true,
		},
		{

			TemplatePath:   path + "/istio/instance/clusterrolebinding.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.ClusterRoleBindingGVR],
			Own:            false,
			Updatable:      true,
		},
		{

			TemplatePath:   path + "/istio/instance/dep.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.DeploymentGVR],
			Own:            false,
			Updatable:      true,
		},
		{

			TemplatePath:   path + "/istio/instance/sa.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SaGVR],
			Own:            false,
			Updatable:      true,
		},
		{

			TemplatePath:   path + "/istio/instance/svc.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.SvcGVR],
			Own:            false,
			Updatable:      true,
		},
		{

			TemplatePath:   path + "/istio/instance/instance.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IstioGVR],
			Own:            false,
			Updatable:      true,
		},
	}
}

func ingressState() []*desired.State {
	return []*desired.State{
		{
			TemplatePath:   path + "/ingress/gw.tpl",
			Template:       nil,
			ParsedTemplate: "",
			Obj:            &unstructured.Unstructured{},
			GVR:            desired.Kinds[desired.IstioGwGVR],
			Own:            true,
			Updatable:      true,
		},
	}
}

func IstioInstanceState(cnvrgInfra *mlopsv1.CnvrgInfra) []*desired.State {
	var state []*desired.State

	if *cnvrgInfra.Spec.Networking.Istio.Enabled {
		state = append(state, istioInstanceState()...)
	}

	if cnvrgInfra.Spec.Networking.Ingress.Type == mlopsv1.IstioIngress && *cnvrgInfra.Spec.Networking.Ingress.IstioGwEnabled {
		state = append(state, ingressState()...)
	}

	return state
}

func CnvrgAppNetworkingState(cnvrgApp *mlopsv1.CnvrgApp) []*desired.State {
	var state []*desired.State

	if cnvrgApp.Spec.Networking.Ingress.Type == mlopsv1.IstioIngress && *cnvrgApp.Spec.Networking.Ingress.IstioGwEnabled {
		state = append(state, ingressState()...)
	}

	return state
}

func IstioCrds() (crds []*desired.State) {
	err := pkger.Walk(path+"/istio/crds", func(path string, info os.FileInfo, err error) error {
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
			Updatable:      false,
		}
		crds = append(crds, crd)
		return nil
	})
	if err != nil {
		zap.S().Error(err, "error loading istio crds")
	}
	return
}
