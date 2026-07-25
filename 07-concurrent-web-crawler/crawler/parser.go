package crawler

import (
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

func Parse(job Job) (Result, error) {
	start := time.Now()

	data, err := os.ReadFile(job.FilePath)
	if err != nil {
		return Result{}, err
	}

	content := string(data)

	info, err := os.Stat(job.FilePath)
	if err != nil {
		return Result{}, err
	}

	result := Result{
		JobID:      job.ID,
		FileName:   filepath.Base(job.FilePath),
		Lines:      countLines(content),
		Words:      countWords(content),
		Characters: countCharacters(content),
		Size:       info.Size(),
		Duration:   time.Since(start),
	}

	return result, nil
}

func countLines(content string) int {
	if content == "" {
		return 0
	}

	return len(strings.Split(content, "\n"))
}

func countWords(content string) int {
	return len(strings.Fields(content))
}

func countCharacters(content string) int {
	return utf8.RuneCountInString(content)
}