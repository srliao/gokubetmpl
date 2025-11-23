package v1

import (
	"maps"

	esv1 "github.com/external-secrets/external-secrets/apis/externalsecrets/v1"
	"github.com/srliao/gokubetmpl/model"
	"github.com/srliao/gokubetmpl/util"
)

type ExternalSecretWrapperV1 struct {
	*esv1.ExternalSecret
}

var _ model.Resource = &ExternalSecretWrapperV1{}

func (e *ExternalSecretWrapperV1) Validate() error { return nil }
func (e *ExternalSecretWrapperV1) Marshal() ([]byte, error) {
	return util.MarshalYamlRemoveEmpty(e.ExternalSecret)
}
func (e *ExternalSecretWrapperV1) MergeAnnotations(annotations map[string]string) *ExternalSecretWrapperV1 {
	if e.Annotations == nil {
		e.Annotations = make(map[string]string)
	}
	maps.Copy(e.Annotations, annotations)
	return e
}
func (e *ExternalSecretWrapperV1) WithAnnotations(annotations map[string]string) *ExternalSecretWrapperV1 {
	e.Annotations = annotations
	return e
}
func (e *ExternalSecretWrapperV1) init() {
	if e.Spec.Target.Template == nil {
		e.Spec.Target.Template = &esv1.ExternalSecretTemplate{
			EngineVersion: esv1.TemplateEngineV2,
		}
	}
	if e.Spec.Target.Template.Data == nil {
		e.Spec.Target.Template.Data = make(map[string]string)
	}
}

func (e *ExternalSecretWrapperV1) AddDataToTemplate(name, from string) *ExternalSecretWrapperV1 {
	e.init()
	e.Spec.Target.Template.Data[name] = from
	return e
}

func (e *ExternalSecretWrapperV1) AddMapDataToTemplate(d map[string]string) *ExternalSecretWrapperV1 {
	e.init()
	maps.Copy(e.Spec.Target.Template.Data, d)
	return e
}

func (e *ExternalSecretWrapperV1) AddExternalDataFromKeyExtract(key string) *ExternalSecretWrapperV1 {
	e.Spec.DataFrom = append(e.Spec.DataFrom, esv1.ExternalSecretDataFromRemoteRef{
		Extract: &esv1.ExternalSecretDataRemoteRef{
			Key: key,
		},
	})
	return e
}
