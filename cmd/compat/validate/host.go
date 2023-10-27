package validate

import (
	"fmt"
	"runtime"

	"compat/pkg/scanner"
	"compat/pkg/scanner/linux"
	pkgvalidate "compat/pkg/validate"
	"compat/schema"
)

const (
	HostCmdName = "validate-host"
)

func RunHostCmd(args []string) error {
	var scanners scanner.Initialized

	if len(args) < 1 {
		return fmt.Errorf("image compatibility schema file path is required")
	}

	ic, err := schema.Load(args[0])
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case scanner.Linux.String():
		scanners, err = linux.Init()
		if err != nil {
			return err
		}
	case scanner.Illumos.String():
		return fmt.Errorf("not implemented")
	case scanner.Windows.String():
		return fmt.Errorf("not implemented")
	default:
		return fmt.Errorf("not supported OS: %s", runtime.GOOS)
	}

	res, err := pkgvalidate.RunHostValidation(ic, scanners)
	if err != nil {
		return err
	}

	fmt.Println("Validation results:")
	for group, status := range res {
		fmt.Printf("%s: %s\n", group, status)
	}

	return nil
}
