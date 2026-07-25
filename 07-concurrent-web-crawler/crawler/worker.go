package crawler

import "sync"

func Worker(
	id int, 
	jobs <-chan Job, 
	results chan<- Result, 
	errors chan<- error, 
	wg *sync.WaitGroup){
		
		defer wg.Done()
		for job := range jobs {
			result, err := Parse(job)
			if err != nil {
				errors <- err
				continue
			}
			
			results <- result

		}
}