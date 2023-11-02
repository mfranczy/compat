package linux

import (
	"compat/pkg/scanner"
)

func Init() (scanner.Factory, error) {
	initialized := make(scanner.Factory)

	scanners := map[string]scanner.InitFunc{
		KernelConfigurationScanner: NewKernelConfiguration,
		KernelCmdlineScanner:       NewKernelCmdline,
		KernelModulesScanner:       NewKernelModules,
		KernelDriversScanner:       NewKernelDrivers,
	}

	for k, sf := range scanners {
		s, err := sf()
		if err != nil {
			return nil, err
		}
		initialized[k] = s
	}

	return initialized, nil
}
