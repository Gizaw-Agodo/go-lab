package crawler

import "time"

type Result struct {
	JobID int 
	FileName string
	Lines int 
	Words int 
	Characters int 
	Size int64
	Duration time.Duration
}