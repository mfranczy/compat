package validate

import (
	"fmt"
	"regexp"

	"github.com/moby/moby/pkg/parsers/kernel"

	"compat/pkg/scanner"
	"compat/pkg/types/v1"
)

type GroupStatus string

const (
	GroupSucceeded GroupStatus = "Succeeded"
	GroupFailed    GroupStatus = "Failed"
)

type Result map[string]GroupStatus

func RunHostValidation(ic *v1.ImageCompatibilitySchema, sc scanner.Factory) (Result, error) {
	kernelVersion, err := matchKernelVersion(ic)
	if err != nil {
		return nil, err
	}
	return validateHost(ic, kernelVersion, sc), nil
}

func matchKernelVersion(ic *v1.ImageCompatibilitySchema) (string, error) {
	kv, err := kernel.GetKernelVersion()
	if err != nil {
		return "", err
	}

	for expr := range ic.Linux {
		r, err := regexp.Compile(expr)
		if err != nil {
			return "", err
		}

		if r.MatchString(kv.String()) {
			return expr, nil
		}
	}

	return "", fmt.Errorf("could not find compatible kernel version")
}

func validateHost(ic *v1.ImageCompatibilitySchema, kernelVersion string, sc scanner.Factory) Result {
	var (
		res = make(Result)
		err error
	)

	for group, schema := range ic.Linux[kernelVersion] {
		if schema.OneOf != nil {
			subGroupsCnt := len(schema.OneOf)
			failedSubGroupsCnt := 0

			for _, subjects := range schema.OneOf {
				if err = runScanners(subjects, sc); err != nil {
					failedSubGroupsCnt++
				}
			}
			if failedSubGroupsCnt >= subGroupsCnt {
				res[group] = GroupFailed
			} else {
				res[group] = GroupSucceeded
			}
		} else {
			if err = runScanners(schema.Subjects, sc); err != nil {
				res[group] = GroupFailed
			} else {
				res[group] = GroupSucceeded
			}
		}
	}

	return res
}

func runScanners(subjects v1.Subjects, sc scanner.Factory) error {
	for subject, input := range subjects {
		if s, ok := sc[subject]; ok {
			if err := s.Run(input); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("could not find scanner for %q subject", subject)
		}
	}
	return nil
}
