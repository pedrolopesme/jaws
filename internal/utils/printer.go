package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/pretty"
)

func Print(title string, content string) error {
	fmt.Println(title)
	formattedContent := pretty.Pretty([]byte(content))
	formattedContent = pretty.Color([]byte(formattedContent), nil)
	fmt.Println(string(formattedContent))
	return nil
}

func PrintTable(header []string, content [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(content)
	table.Render()
}

func PrinBreaklines(numLines int) {
	for i := 0; i < numLines; i++ {
		fmt.Println()
	}
}
