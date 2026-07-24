package main

import (
	"fmt"
	"go-lab/06-concurrent-file-analyzer/analyzer"
	"log"
)

func main() {
	paths := []string{
		"input/books.txt",
		"input/logs.txt",
		"input/notes.txt",
		"input/article.txt",
	}

	results, errs := analyzer.AnalyzeFiles(paths)

	if len(results) == 0 && len(errs) > 0 {
		log.Fatal("no files were analyzed successfully")
	}

	report := analyzer.BuildReport(results, errs)

	fmt.Println(report)
}