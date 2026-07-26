package crawler

import (
	"sync"
	"time"
)

type Statistics struct {
	mu sync.Mutex
	ProcessedJobs int 
	SuccessfullJobs int 
	FailedJobs int 

	TotalBytes int64
	Duration time.Duration
}

type StatisticsSnapshot struct {
	ProcessedJobs int
	SuccessfulJobs int
	FailedJobs int

	TotalBytes int64
	TotalDuration time.Duration
}

func (s *Statistics) RecordSuccess(result Result){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ProcessedJobs ++ 
	s.SuccessfullJobs ++ 
	s.TotalBytes += result.Size
	s.Duration += result.Duration
}

func (s *Statistics) RecordFailure(){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ProcessedJobs ++
	s.FailedJobs ++

}

func (s *Statistics) Snapshot() StatisticsSnapshot {
	s.mu.Lock()
	defer s.mu.Unlock()

	return StatisticsSnapshot{
		ProcessedJobs: s.ProcessedJobs,
		SuccessfulJobs: s.SuccessfullJobs,
		FailedJobs: s.FailedJobs,
		TotalBytes: s.TotalBytes,
		TotalDuration: s.Duration,
	}
}