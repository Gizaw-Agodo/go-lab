package analyzer

import (
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

func AnalyzeFile(path string) (Result, error){
	start := time.Now()
	
	// read the file 
	data, err := os.ReadFile(path)
	if err != nil {
		return Result{}, err
	}

	content := string(data)
	info, err := os.Stat(path)
	if err != nil {
		return Result{}, err 
	}
	lines := countLines(content)
	words := countWords(content)
	characters := countCharacters(content)

	// build the result and return 
	result := Result{
		Filename: filepath.Base(path),
		Lines: lines,
		Words: words,
		Characters: characters,
		Size: info.Size(),
		Duration: time.Since(start),
	}

	return result,nil
}

func countLines(content string ) int {
	 if content == "" {
		return 0
	 }
	 lines := strings.Split(content, "\n")
	 return len(lines)
}

func countWords(content string) int {
	if content == "" {
		return 0
	}
	fields := strings.Fields(content)
	return len(fields)
}

func countCharacters(content string ) int {
	return utf8.RuneCountInString(content)
}