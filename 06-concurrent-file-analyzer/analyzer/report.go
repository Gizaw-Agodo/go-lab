package analyzer

import (
	"fmt"
	"strings"
)

func BuildReport(results []Result, errs []error) string {
	var builder strings.Builder

	builder.WriteString("=========================================\n")
	builder.WriteString("     Concurrent File Analyzer Report\n")
	builder.WriteString("=========================================\n\n")

	builder.WriteString(fmt.Sprintf("Files Analyzed : %d\n", len(results)))
	builder.WriteString(fmt.Sprintf("Errors         : %d\n\n", len(errs)))

	builder.WriteString("-----------------------------------------\n")

	for _, result := range results {
		builder.WriteString(fmt.Sprintf("File       : %s\n", result.Filename))
		builder.WriteString(fmt.Sprintf("Lines      : %d\n", result.Lines))
		builder.WriteString(fmt.Sprintf("Words      : %d\n", result.Words))
		builder.WriteString(fmt.Sprintf("Characters : %d\n", result.Characters))
		builder.WriteString(fmt.Sprintf("Size       : %d bytes\n", result.Size))
		builder.WriteString(fmt.Sprintf("Duration   : %v\n", result.Duration))
		builder.WriteString("-----------------------------------------\n")
	}

	totalLines, totalWords, totalCharacters, totalSize := calculateTotals(results)

	builder.WriteString("\n========== TOTALS ==========\n")
	builder.WriteString(fmt.Sprintf("Total Files      : %d\n", len(results)))
	builder.WriteString(fmt.Sprintf("Total Lines      : %d\n", totalLines))
	builder.WriteString(fmt.Sprintf("Total Words      : %d\n", totalWords))
	builder.WriteString(fmt.Sprintf("Total Characters : %d\n", totalCharacters))
	builder.WriteString(fmt.Sprintf("Total Size       : %d bytes\n", totalSize))

	if len(errs) > 0 {
		builder.WriteString("\n========== ERRORS ==========\n")

		for _, err := range errs {
			builder.WriteString(fmt.Sprintf("- %v\n", err))
		}
	}

	return builder.String()
}

func calculateTotals(results []Result) (lines, words, characters int, size int64) {
	for _, result := range results {
		lines += result.Lines
		words += result.Words
		characters += result.Characters
		size += result.Size
	}

	return
}