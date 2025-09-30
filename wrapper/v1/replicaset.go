package v1

import (
	"github.com/srliao/gokubetmpl/model"
	"github.com/srliao/gokubetmpl/util"
	appsv1 "k8s.io/api/apps/v1"
)

type ReplicaSetWrapper struct {
	*appsv1.ReplicaSet
}

var _ model.Resource = &ReplicaSetWrapper{}

func (r *ReplicaSetWrapper) Validate() error { return nil }
func (r *ReplicaSetWrapper) Marshal() ([]byte, error) {
	return util.MarshalYamlRemoveEmpty(r.ReplicaSet)
}
