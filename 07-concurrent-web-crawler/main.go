package main

import (
	"fmt"
	"log"

	"go-lab/07-concurrent-web-crawler/crawler"
)

func main() {
	paths := []string{
		"pages/page1.txt",
		"pages/page2.txt",
		"pages/page3.txt",
		"pages/page4.txt",
		"pages/page5.txt",
	}

	const workerCount = 3

	results, errs := crawler.Crawl(paths, workerCount)

	if len(results) == 0 && len(errs) > 0 {
		log.Fatal("no pages were processed successfully")
	}

	report := crawler.BuildReport(results, errs)

	fmt.Println(report)
}