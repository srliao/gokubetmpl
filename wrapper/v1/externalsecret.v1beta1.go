package v1

import (
	"maps"

	esv1beta1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"github.com/srliao/gokubetmpl/model"
	"github.com/srliao/gokubetmpl/util"
)

type ExternalSecretWrapperV1Beta1 struct {
	*esv1beta1.ExternalSecret
}

var _ model.Resource = &ExternalSecretWrapperV1Beta1{}

func (e *ExternalSecretWrapperV1Beta1) Validate() error { return nil }
func (e *ExternalSecretWrapperV1Beta1) Marshal() ([]byte, error) {
	return util.MarshalYamlRemoveEmpty(e.ExternalSecret)
}
func (e *ExternalSecretWrapperV1Beta1) MergeAnnotations(annotations map[string]string) *ExternalSecretWrapperV1Beta1 {
	if e.Annotations == nil {
		e.Annotations = make(map[string]string)
	}
	maps.Copy(e.Annotations, annotations)
	return e
}
func (e *ExternalSecretWrapperV1Beta1) WithAnnotations(annotations map[string]string) *ExternalSecretWrapperV1Beta1 {
	e.Annotations = annotations
	return e
}
func (e *ExternalSecretWrapperV1Beta1) init() {
	if e.Spec.Target.Template == nil {
		e.Spec.Target.Template = &esv1beta1.ExternalSecretTemplate{
			EngineVersion: esv1beta1.TemplateEngineV2,
		}
	}
	if e.Spec.Target.Template.Data == nil {
		e.Spec.Target.Template.Data = make(map[string]string)
	}
}

func (e *ExternalSecretWrapperV1Beta1) AddDataToTemplate(name, from string) *ExternalSecretWrapperV1Beta1 {
	e.init()
	e.Spec.Target.Template.Data[name] = from
	return e
}

func (e *ExternalSecretWrapperV1Beta1) AddMapDataToTemplate(d map[string]string) *ExternalSecretWrapperV1Beta1 {
	e.init()
	maps.Copy(e.Spec.Target.Template.Data, d)
	return e
}

func (e *ExternalSecretWrapperV1Beta1) AddExternalDataFromKeyExtract(key string) *ExternalSecretWrapperV1Beta1 {
	e.Spec.DataFrom = append(e.Spec.DataFrom, esv1beta1.ExternalSecretDataFromRemoteRef{
		Extract: &esv1beta1.ExternalSecretDataRemoteRef{
			Key: key,
		},
	})
	return e
}
