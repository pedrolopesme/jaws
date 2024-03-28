package utils

import (
	"encoding/json"
	"fmt"
)

func Print(title string, output string, color func(format string, a ...interface{})) error {
	color("\n%v:", title)

	var data interface{}
	err := json.Unmarshal([]byte(output), &data)
	if err != nil {
		return err
	}

	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(formattedJSON))
	return nil
}
