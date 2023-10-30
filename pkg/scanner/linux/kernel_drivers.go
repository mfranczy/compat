package linux

import (
	"fmt"
	"github.com/mfranczy/compat/pkg/scanner"
)

const KernelDriversScanner = "kernelDrivers"

type KernelDrivers struct {
	id    string
	group string
	input interface{}
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

func (d *KernelDrivers) Input() interface{} {
	return d.input
}

func (d *KernelDrivers) Run() error {
	return fmt.Errorf("test")
}

func NewKernelDrivers(id string, group string, input interface{}) (scanner.Scanner, error) {
	return &KernelDrivers{
		id:    id,
		group: group,
		input: input,
	}, nil
}

func init() {
	scanner.Register(scanner.Linux, KernelDriversScanner, NewKernelDrivers)
}
