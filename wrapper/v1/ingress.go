package v1

import (
	"maps"

	"github.com/srliao/gokubetmpl/model"
	"github.com/srliao/gokubetmpl/util"
	networkingv1 "k8s.io/api/networking/v1"
)

type IngressWrapper struct {
	*networkingv1.Ingress
}

var _ model.KubeResource = &IngressWrapper{}
var _ model.Annotatable[*IngressWrapper] = &IngressWrapper{}

func (i *IngressWrapper) Validate() error          { return nil }
func (i *IngressWrapper) Marshal() ([]byte, error) { return util.MarshalYamlRemoveEmpty(i.Ingress) }

func (i *IngressWrapper) MergeAnnotations(annotations map[string]string) *IngressWrapper {
	if i.Annotations == nil {
		i.Annotations = make(map[string]string)
	}
	maps.Copy(i.Annotations, annotations)
	return i
}

func (i *IngressWrapper) WithAnnotations(annotations map[string]string) *IngressWrapper {
	i.Annotations = annotations
	return i
}

func (i *IngressWrapper) WithRules(rules []networkingv1.IngressRule) *IngressWrapper {
	i.Spec.Rules = rules
	return i
}

func (i *IngressWrapper) AddRule(rule networkingv1.IngressRule) *IngressWrapper {
	i.Spec.Rules = append(i.Spec.Rules, rule)
	return i
}

func (i *IngressWrapper) WithTLS(tls []networkingv1.IngressTLS) *IngressWrapper {
	i.Spec.TLS = tls
	return i
}

func (i *IngressWrapper) WithClass(class string) *IngressWrapper {
	i.Spec.IngressClassName = &class
	return i
}
