package analyzer

import "sync"


func AnalyzeFiles(paths []string) ([]Result, []error) {
	results := make(chan Result, len(paths))
	errors := make(chan error, len(paths))

	var wg sync.WaitGroup
	wg.Add(len(paths))

	for _, path := range paths {
		go Worker(path, &wg, results, errors)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	var allResults []Result
	var allErrors []error

	for results != nil || errors != nil {
		select {
		case result, ok := <-results:
			if !ok {
				results = nil
				continue
			}
			allResults = append(allResults, result)

		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}
			allErrors = append(allErrors, err)
		}
	}

	return allResults, allErrors
}