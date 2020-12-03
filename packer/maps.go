package packer

import (
	"fmt"

	packersdk "github.com/hashicorp/packer/packer-plugin-sdk/packer"
)

type MapOfProvisioner map[string]func() (packersdk.Provisioner, error)

func (mop MapOfProvisioner) Has(provisioner string) bool {
	_, res := mop[provisioner]
	return res
}

func (mop MapOfProvisioner) Start(provisioner string) (packersdk.Provisioner, error) {
	p, found := mop[provisioner]
	if !found {
		return nil, fmt.Errorf("Unknown provisioner %s", provisioner)
	}
	return p()
}

func (mop MapOfProvisioner) List() []string {
	res := []string{}
	for k := range mop {
		res = append(res, k)
	}
	return res
}

type MapOfPostProcessor map[string]func() (packersdk.PostProcessor, error)

func (mopp MapOfPostProcessor) Has(postProcessor string) bool {
	_, res := mopp[postProcessor]
	return res
}

func (mopp MapOfPostProcessor) Start(postProcessor string) (packersdk.PostProcessor, error) {
	p, found := mopp[postProcessor]
	if !found {
		return nil, fmt.Errorf("Unknown post-processor %s", postProcessor)
	}
	return p()
}

func (mopp MapOfPostProcessor) List() []string {
	res := []string{}
	for k := range mopp {
		res = append(res, k)
	}
	return res
}

type MapOfBuilder map[string]func() (packersdk.Builder, error)

func (mob MapOfBuilder) Has(builder string) bool {
	_, res := mob[builder]
	return res
}

func (mob MapOfBuilder) Start(builder string) (packersdk.Builder, error) {
	d, found := mob[builder]
	if !found {
		return nil, fmt.Errorf("Unknown builder %s", builder)
	}
	return d()
}

func (mob MapOfBuilder) List() []string {
	res := []string{}
	for k := range mob {
		res = append(res, k)
	}
	return res
}
