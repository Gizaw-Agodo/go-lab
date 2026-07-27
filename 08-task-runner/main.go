package main

import (
	"context"
	"go-lab/08-task-runner/runner"
	"time"
)

func main(){
	tasks := []runner.Task{
		{ID: 1, Name: "Download File"},
		{ID: 2, Name: "Generate Report"},
		{ID: 3, Name: "Send Email"},
	}

	r := runner.NewRunner(tasks)
	cxt,cancel := context.WithCancel(context.Background())
	defer cancel()
	
	go func (){
		time.Sleep(5*time.Second)
		cancel()
	}()
	
	r.Run(cxt)

}
