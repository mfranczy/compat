package linux

import (
	"github.com/mfranczy/compat/pkg/types/v1/hardware"
)

const (
	KernelConfigurationSchemaName = "kernelConfiguration"
	KernelCmdlineSchemaName       = "kernelCmdline"
	KernelModulesSchemaName       = "kernelModules"
	KernelDriversSchemaName       = "kernelDrivers"
)

type Schema struct {
	OneOf map[string]*Configuration `json:"oneof,omitempty"`
	Configuration
}

type Configuration struct {
	Hardware            *hardware.Schema     `json:"hardware,omitempty"`
	KernelConfiguration *KernelConfiguration `json:"kernelConfiguration,omitempty"`
	KernelCmdline       *KernelCmdline       `json:"kernelCmdline,omitempty"`
	KernelModules       *KernelModules       `json:"kernelModules,omitempty"`
	KernelDrivers       *KernelDrivers       `json:"kernelDrivers,omitempty"`
}

type KernelConfiguration struct {
}

func (KernelConfiguration) String() string {
	return KernelConfigurationSchemaName
}

type KernelCmdline struct {
}

func (KernelCmdline) String() string {
	return KernelCmdlineSchemaName
}

type KernelModules struct {
}

func (KernelModules) String() string {
	return KernelModulesSchemaName
}

type KernelDrivers struct {
}

func (KernelDrivers) String() string {
	return KernelDriversSchemaName
}
