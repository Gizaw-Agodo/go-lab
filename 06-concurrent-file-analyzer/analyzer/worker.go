package analyzer

import "sync"


func Worker(path string, wg *sync.WaitGroup, ch chan<- Result, errors chan<- error) {
	defer wg.Done()
	result,err := AnalyzeFile(path)
	if err != nil {
		errors<- err 
		return 
	}
	ch <- result 
}

