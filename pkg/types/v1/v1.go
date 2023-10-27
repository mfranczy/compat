package v1

import (
	"compat/pkg/types/v1/linux"
)

type ImageCompatibilitySchema struct {
	Version string      `json:"schemaVersion"`
	Linux   LinuxSchema `json:"linux,omitempty"`
}

type LinuxSchema map[string]map[string]*linux.Schema
