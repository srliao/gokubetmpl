package model

import "sigs.k8s.io/controller-runtime/pkg/client"

type Resource interface {
	client.Object
	Validate() error
	Marshal() ([]byte, error)
}

type File interface {
	Name() string
	Content() []byte
}

type Annotatable[K any] interface {
	MergeAnnotations(annotations map[string]string) K
	WithAnnotations(annotations map[string]string) K
}
