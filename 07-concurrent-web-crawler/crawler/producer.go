package crawler

func Produce(
	paths []string,
	jobs chan<- Job,
){
	defer close(jobs)
	
	for i, path := range paths {
		job := Job{
			ID : i + 1, 
			FilePath: path,
		}
		jobs <- job 
	}

}