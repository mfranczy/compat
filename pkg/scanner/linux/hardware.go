package linux

import (
	"github.com/mfranczy/compat/pkg/scanner"
	"github.com/mfranczy/compat/pkg/types/v1/hardware"
)

const HardwareScanner = hardware.SchemaName

type Hardware struct {
	id    string
	group string
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

func (h *Hardware) Run(i interface{}) error {
	return nil
}

func NewHardware(id string, group string) scanner.Scanner {
	return &Hardware{
		id:    id,
		group: group,
	}
}

func init() {
	scanner.Register(scanner.Linux, HardwareScanner, NewHardware)
}
