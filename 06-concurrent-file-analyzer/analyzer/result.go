package analyzer
import "time"

type Result struct {
	Filename string 
	Lines int 
	Words int 
	Characters int 
	Size int64
	Duration time.Duration
}

