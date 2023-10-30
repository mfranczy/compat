package validate

import (
	"fmt"
	"github.com/mfranczy/compat/pkg/types/v1/linux"
	"regexp"
	"runtime"

	"github.com/mfranczy/compat/pkg/scanner"
	v1 "github.com/mfranczy/compat/pkg/types/v1"
	"github.com/mfranczy/compat/pkg/worker"
	"github.com/moby/moby/pkg/parsers/kernel"
)

func Run(ic *v1.ImageCompatibilitySchema, sc scanner.OsScanners, workersNum int) (map[string]map[string]bool, error) {
	result := make(map[string]map[string]bool)
	p := worker.NewPool(workersNum)
	go p.Start()

	switch runtime.GOOS {
	case scanner.Linux.String():
		kv, err := kernel.GetKernelVersion()
		if err != nil {
			return nil, err
		}

		var validKernel string
		for expr := range ic.Linux {
			r, err := regexp.Compile(expr)
			if err != nil {
				return nil, err
			}

			if r.MatchString(kv.String()) {
				validKernel = expr
				break
			}
		}

		if validKernel == "" {
			return nil, fmt.Errorf("could not find compatible kernel version")
		}

		// TODO: remove parallelism
		for id, schema := range ic.Linux[validKernel] {
			if schema.OneOf != nil {
				groupName := id
				for id, subjects := range schema.OneOf {
					job, err := createJob(id, groupName, subjects)
					if err != nil {
						return nil, err
					}
					p.Jobs <- job
				}
			} else {
				job, err := createJob(id, id, schema.Subjects)
				if err != nil {
					return nil, err
				}
				p.Jobs <- job
			}
		}
		close(p.Jobs)

		// TODO: fix this, it's very dummy
		for job := range p.Results {
			if _, ok := result[job.SubjectGroup]; !ok {
				result[job.SubjectGroup] = make(map[string]bool)
			}
			for _, task := range job.Tasks {
				result[job.SubjectGroup][task.ID()] = job.Succeeded()
			}
		}
		break
	case scanner.Illumos.String():
		return nil, fmt.Errorf("not implemented")
	case scanner.Windows.String():
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf("not supported OS: %s", runtime.GOOS)
	}

	return result, nil
}

func createJob(id, groupName string, subjects linux.Subjects) (*worker.Job, error) {
	job := worker.NewJob()
	for subject, input := range subjects {
		if creator, ok := scanner.Scanners.Linux[subject]; ok {
			job.SubjectGroup = groupName
			job.SubjectID = id
			sc, err := creator(id, groupName, input)
			if err != nil {
				return nil, err
			}
			job.Tasks = append(job.Tasks, sc)
		} else {
			return nil, fmt.Errorf("unregistered scanner for %s subject", subject)
		}
		job.Tasks = append(job.Tasks)
	}

	return job, nil
}
