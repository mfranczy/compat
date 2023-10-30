package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
)

const KernelConfigurationScanner = "kernelConfiguration"

type KernelConfiguration struct {
	id    string
	group string
	input interface{}
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

func (c *KernelConfiguration) Input() interface{} {
	return c.input
}

func (c *KernelConfiguration) Run() error {
	return nil
}

func NewKernelConfiguration(id string, group string, input interface{}) (scanner.Scanner, error) {
	return &KernelConfiguration{
		id:    id,
		group: group,
		input: input,
	}, nil
}

func init() {
	scanner.Register(scanner.Linux, KernelConfigurationScanner, NewKernelConfiguration)
}
