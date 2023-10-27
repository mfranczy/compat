package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
	"github.com/mfranczy/compat/pkg/types/v1/linux"
)

const KernelConfigurationScanner = linux.KernelConfigurationSchemaName

type KernelConfiguration struct {
	id    string
	group string
}

func (c *KernelConfiguration) Name() string {
	return KernelConfigurationScanner
}

func (c *KernelConfiguration) ID() string {
	return c.id
}

func (c *KernelConfiguration) Group() string {
	return c.group
}

func (c *KernelConfiguration) Run(i interface{}) error {
	return nil
}

func NewKernelConfiguration(id string, group string) scanner.Scanner {
	return &KernelConfiguration{
		id:    id,
		group: group,
	}
}

func init() {
	scanner.Register(scanner.Linux, KernelConfigurationScanner, NewKernelConfiguration)
}
