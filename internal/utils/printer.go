package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/pretty"
)

// Print prints a title and its content in a visually appealing way.
func Print(title string, content string) {
	fmt.Println(title)
	formattedContent := pretty.Pretty([]byte(content))
	formattedContent = pretty.Color([]byte(formattedContent), nil)
	fmt.Println(string(formattedContent))
	PrinBreaklines(1)
}

// PrintTable prints a table with the given header and content.
func PrintTable(header []string, content [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(content)
	table.Render()
	PrinBreaklines(2)
}

// PrinBreaklines prints a number of newlines.
func PrinBreaklines(numLines int) {
	for i := 0; i < numLines; i++ {
		fmt.Println()
	}
}
