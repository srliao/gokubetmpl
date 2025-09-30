package v1

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

type ConfigMapWrapper struct {
	*corev1.ConfigMap
}

func (c *ConfigMapWrapper) Validate() error          { return nil }
func (c *ConfigMapWrapper) Marshal() ([]byte, error) { return yaml.Marshal(c.ConfigMap) }
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
