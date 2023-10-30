package worker

import (
	"sync"
)

const (
	bufferSize = 10
)

type Pool struct {
	wg         sync.WaitGroup
	workersNum int
	Jobs       chan *Job
	Results    chan *Job
}

func (p *Pool) Start() {
	defer close(p.Results)
	for i := 0; i < p.workersNum; i++ {
		p.wg.Add(1)
		go Run(&p.wg, p.Jobs, p.Results)
	}
	p.wg.Wait()
}

func NewPool(workersNum int) *Pool {
	return &Pool{
		workersNum: workersNum,
		Jobs:       make(chan *Job, bufferSize),
		Results:    make(chan *Job, bufferSize),
	}
}
