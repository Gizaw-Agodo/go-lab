package runner

import (
	"context"
	"sync"
)

type Runner struct {
	Tasks []Task
}

func NewRunner(tasks []Task ) *Runner {
	runner := Runner{
		Tasks:tasks,
	}
	return &runner
}

func (r *Runner) Run(cxt context.Context){
	var wg sync.WaitGroup

	for _, task := range r.Tasks {
		wg.Add(1)
		go func(task Task){
			defer wg.Done()
			Worker(cxt, task)
		}(task)
	}
	wg.Wait()
}