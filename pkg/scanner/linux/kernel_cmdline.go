package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
	"github.com/mfranczy/compat/pkg/types/v1/linux"
)

const KernelCmdlineScanner = linux.KernelCmdlineSchemaName

type KernelCmdline struct {
	id    string
	group string
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

func (c *KernelCmdline) Run(i interface{}) error {
	return nil
}

func NewKernelCmdline(id string, group string) scanner.Scanner {
	return &KernelCmdline{
		id:    id,
		group: group,
	}
}

func init() {
	scanner.Register(scanner.Linux, KernelCmdlineScanner, NewKernelCmdline)
}
