package hardware

import "strings"

const (
	SchemaName string = "hardware"
)

type Schema struct {
	CPU []string `json:"cpu,omitempty"`
	PCI []PCI    `json:"pci,omitempty"`
	USB []USB    `json:"usb,omitempty"`
}

func (s *Schema) String() string {
	return SchemaName
}

type PCI string

func (p PCI) String() string {
	return string(p)
}

func (p PCI) ClassID() string {
	return getRegister(p.String(), 0)
}

func (p PCI) VendorID() string {
	return getRegister(p.String(), 1)
}

func (p PCI) ProductID() string {
	return getRegister(p.String(), 2)
}

type USB string

func (u USB) String() string {
	return string(u)
}

func (u USB) ClassID() string {
	return getRegister(u.String(), 0)
}

func (u USB) VendorID() string {
	return getRegister(u.String(), 1)
}

func (u USB) ProductID() string {
	return getRegister(u.String(), 2)
}

// getRegister returns:
// - Class ID at index 0
// - Vendor ID at index 1
// - Product ID at index 0
// for PCI and USB devices
func getRegister(s string, i int) string {
	o := strings.Split(s, ":")
	if len(o) >= i {
		return o[i]
	}
	return ""
}
