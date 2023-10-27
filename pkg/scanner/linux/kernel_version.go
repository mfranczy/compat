package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
)

const KernelVersionScanner = "kernelVersion"

type KernelVersion struct {
	id    string
	group string
}

func (d *KernelVersion) Name() string {
	return KernelVersionScanner
}

func (d *KernelVersion) ID() string {
	return d.id
}

func (d *KernelVersion) Group() string {
	return d.group
}

func (d *KernelVersion) Run(i interface{}) error {
	return nil
}

func NewKernelVersion(id string, group string) scanner.Scanner {
	return &KernelVersion{
		id:    id,
		group: group,
	}
}

func init() {
	scanner.Register(scanner.Linux, KernelVersionScanner, NewKernelVersion)
}
