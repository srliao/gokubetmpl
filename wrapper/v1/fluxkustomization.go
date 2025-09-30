package v1

import (
	"path/filepath"

	fluxv1 "github.com/fluxcd/kustomize-controller/api/v1"
	"github.com/srliao/gokubetmpl/model"
	"github.com/srliao/gokubetmpl/util"
)

type FluxKustomizationWrapper struct {
	*fluxv1.Kustomization
	subpath string
}

var _ model.Resource = &FluxKustomizationWrapper{}

func (f *FluxKustomizationWrapper) Marshal() ([]byte, error) {
	return util.MarshalYamlRemoveEmpty(f.Kustomization)
}
func (f *FluxKustomizationWrapper) Validate() error { return nil }
func (f *FluxKustomizationWrapper) WithPath(base, subpath string) *FluxKustomizationWrapper {
	f.Kustomization.Spec.Path = filepath.Join(base, subpath)
	f.subpath = subpath
	return f
}
func (f *FluxKustomizationWrapper) WithDependsOn(deps []fluxv1.DependencyReference) *FluxKustomizationWrapper {
	f.Kustomization.Spec.DependsOn = deps
	return f
}
