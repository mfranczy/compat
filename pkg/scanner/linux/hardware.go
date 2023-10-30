package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
)

const HardwareScanner = "hardware"

type Hardware struct {
	id    string
	group string
	input interface{}
}

func (h *Hardware) Name() string {
	return HardwareScanner
}

func (h *Hardware) ID() string {
	return h.id
}

func (h *Hardware) Group() string {
	return h.group
}

func (h *Hardware) Input() interface{} {
	return h.input
}

func (h *Hardware) Run() error {
	return nil
}

func NewHardware(id string, group string, input interface{}) (scanner.Scanner, error) {
	return &Hardware{
		id:    id,
		group: group,
		input: input,
	}, nil
}

func init() {
	scanner.Register(scanner.Linux, HardwareScanner, NewHardware)
}
