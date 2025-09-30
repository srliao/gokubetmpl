package v1

import (
	vs1alpha1 "github.com/backube/volsync/api/v1alpha1"
	"github.com/srliao/gokubetmpl/model"
	"github.com/srliao/gokubetmpl/util"
)

type ReplicationDestinationWrapper struct {
	*vs1alpha1.ReplicationDestination
}

type ReplicationSourceWrapper struct {
	*vs1alpha1.ReplicationSource
}

var _ model.KubeResource = &ReplicationDestinationWrapper{}
var _ model.KubeResource = &ReplicationSourceWrapper{}

func (r *ReplicationDestinationWrapper) Validate() error { return nil }
func (r *ReplicationDestinationWrapper) Marshal() ([]byte, error) {
	return util.MarshalYamlRemoveEmpty(r.ReplicationDestination)
}

func (r *ReplicationSourceWrapper) Validate() error { return nil }
func (r *ReplicationSourceWrapper) Marshal() ([]byte, error) {
	return util.MarshalYamlRemoveEmpty(r.ReplicationSource)
}
