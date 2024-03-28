package utils

import (
	"fmt"

	"github.com/tidwall/pretty"
)

func Print(title string, content string) error {
	fmt.Println(title)
	formattedContent := pretty.Pretty([]byte(content))
	formattedContent = pretty.Color([]byte(formattedContent), nil)
	fmt.Println(string(formattedContent))
	return nil
}
