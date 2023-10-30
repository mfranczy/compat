package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
)

const KernelModulesScanner = "kernelModules"

type KernelModules struct {
	id    string
	group string
	input interface{}
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

func (d *KernelModules) Input() interface{} {
	return d.input
}

func (d *KernelModules) Run() error {
	return nil
}

func NewKernelModules(id string, group string, input interface{}) (scanner.Scanner, error) {
	return &KernelModules{
		id:    id,
		group: group,
		input: input,
	}, nil
}

func init() {
	scanner.Register(scanner.Linux, KernelModulesScanner, NewKernelModules)
}
