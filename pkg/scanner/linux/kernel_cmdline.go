package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
)

const KernelCmdlineScanner = "kernelCmdline"

type KernelCmdline struct {
	id    string
	group string
	input interface{}
}

func (c *KernelCmdline) Name() string {
	return KernelCmdlineScanner
}

func (c *KernelCmdline) ID() string {
	return c.id
}

func (c *KernelCmdline) Group() string {
	return c.group
}

func (c *KernelCmdline) Input() interface{} {
	return c.input
}

func (c *KernelCmdline) Run() error {
	return nil
}

func NewKernelCmdline(id string, group string, input interface{}) (scanner.Scanner, error) {
	return &KernelCmdline{
		id:    id,
		group: group,
		input: input,
	}, nil
}

func init() {
	scanner.Register(scanner.Linux, KernelCmdlineScanner, NewKernelCmdline)
}
