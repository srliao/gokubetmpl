package v1

import (
	"github.com/srliao/gokubetmpl/model"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/yaml"
)

type ServiceWrapper struct {
	*corev1.Service
}

var _ model.KubeResource = &ServiceWrapper{}

func (s *ServiceWrapper) Validate() error          { return nil }
func (s *ServiceWrapper) Marshal() ([]byte, error) { return yaml.Marshal(s.Service) }

func (s *ServiceWrapper) WithSelector(labels map[string]string) *ServiceWrapper {
	s.Spec.Selector = labels
	return s
}

func (s *ServiceWrapper) AddPort(port, targetPort int32, protocol corev1.Protocol) *ServiceWrapper {
	s.Spec.Ports = append(s.Spec.Ports, corev1.ServicePort{
		Port:       port,
		TargetPort: intstr.FromInt(int(targetPort)),
		Protocol:   protocol,
	})
	return s
}

func (s *ServiceWrapper) WithServiceType(serviceType corev1.ServiceType) *ServiceWrapper {
	s.Spec.Type = serviceType
	return s
}
