package linux

import (
	"compat/pkg/scanner"
)

func Init() (scanner.Initialized, error) {
	initialized := make(scanner.Initialized)

	scanners := map[string]scanner.InitFunc{
		KernelConfigurationScanner: NewKernelConfiguration,
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
