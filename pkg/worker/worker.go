package worker

import (
	"sync"

	"github.com/mfranczy/compat/pkg/scanner"

	"github.com/lithammer/shortuuid/v4"
)

func Run(wg *sync.WaitGroup, jobs <-chan *Job, results chan<- *Job) {
	defer wg.Done()
	for job := range jobs {
		for _, task := range job.Tasks {
			if err := task.Run(); err != nil {
				job.Failed = append(job.Failed, task.ID())
			} else {
				job.Successful = append(job.Successful, task.ID())
			}
		}
		results <- job
	}
}

type Job struct {
	id           string //cpu
	SubjectGroup string // intel or amd
	SubjectID    string
	Tasks        []scanner.Scanner //kernelCmdline, kernelConfiguration...
	Successful   []string
	Failed       []string
}

func (j *Job) ID() string {
	return j.id
}

func (j *Job) Succeeded() bool {
	return len(j.Failed) < 1
}

func NewJob() *Job {
	return &Job{id: shortuuid.New()}
}
