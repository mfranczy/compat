package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
	"github.com/mfranczy/compat/pkg/types/v1/linux"
)

const KernelModulesScanner = linux.KernelModulesSchemaName

type KernelModules struct {
	id    string
	group string
}

func (d *KernelModules) Name() string {
	return KernelModulesScanner
}

func (d *KernelModules) ID() string {
	return d.id
}

func (d *KernelModules) Group() string {
	return d.group
}

func (d *KernelModules) Run(i interface{}) error {
	return nil
}

func NewKernelModules(id string, group string) scanner.Scanner {
	return &KernelModules{
		id:    id,
		group: group,
	}
}

func init() {
	scanner.Register(scanner.Linux, KernelModulesScanner, NewKernelModules)
}
