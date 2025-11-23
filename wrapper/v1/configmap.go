package v1

import (
	"maps"

	"github.com/srliao/gokubetmpl/model"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

type ConfigMapWrapper struct {
	*corev1.ConfigMap
}

var _ model.Resource = &ConfigMapWrapper{}
var _ model.Annotatable[*ConfigMapWrapper] = &ConfigMapWrapper{}

func (c *ConfigMapWrapper) Validate() error          { return nil }
func (c *ConfigMapWrapper) Marshal() ([]byte, error) { return yaml.Marshal(c.ConfigMap) }
func (c *ConfigMapWrapper) MergeAnnotations(annotations map[string]string) *ConfigMapWrapper {
	if c.Annotations == nil {
		c.Annotations = make(map[string]string)
	}
	maps.Copy(c.Annotations, annotations)
	return c
}
func (c *ConfigMapWrapper) WithAnnotations(annotations map[string]string) *ConfigMapWrapper {
	c.Annotations = annotations
	return c
}
func (c *ConfigMapWrapper) WithData(data map[string]string) *ConfigMapWrapper {
	c.Data = data
	return c
}
func (c *ConfigMapWrapper) AddData(key, data string) *ConfigMapWrapper {
	if c.Data == nil {
		c.Data = map[string]string{}
	}
	c.Data[key] = data
	return c
}
