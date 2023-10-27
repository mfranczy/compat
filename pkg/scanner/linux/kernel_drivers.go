package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
	"github.com/mfranczy/compat/pkg/types/v1/linux"
)

const KernelDriversScanner = linux.KernelDriversSchemaName

type KernelDrivers struct {
	id    string
	group string
}

func (d *KernelDrivers) Name() string {
	return KernelDriversScanner
}

func (d *KernelDrivers) ID() string {
	return d.id
}

func (d *KernelDrivers) Group() string {
	return d.group
}

func (d *KernelDrivers) Run(i interface{}) error {
	return nil
}

func NewKernelDrivers(id string, group string) scanner.Scanner {
	return &KernelDrivers{
		id:    id,
		group: group,
	}
}

func init() {
	scanner.Register(scanner.Linux, KernelDriversScanner, NewKernelDrivers)
}
