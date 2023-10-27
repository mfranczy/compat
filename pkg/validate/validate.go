package validate

import (
	"github.com/mfranczy/compat/pkg/scanner"
	v1 "github.com/mfranczy/compat/pkg/types/v1"
)

type Report struct {
}

func Run(ic *v1.ImageCompatibilitySchema, sc scanner.OsScanner) (Report, error) {
	// Run kernel version validation
	// Run all tasks divided to groups (group is a list)
	return Report{}, nil
}
