package utils

import "reflect"

// Print outputs a section of a JWT token, supporting title and a color.
//
//   - title : section name
//   - output: content to be printed
//   - color: color spec from fatih/color package
func Print(title string, output map[string]interface{}, color func(format string, a ...interface{})) {
	color("\n%v:", title)
	for key, val := range output {
		PrintLine(key, val, color)
	}
}

// PrintLine: Output lines, formatting according to its content.
//
//   - key : claim name
//   - val: claim content
//   - color: color spec from fatih/color package
func PrintLine(key string, val interface{}, color func(format string, a ...interface{})) {
	switch val.(type) {
	case []interface{}:
		color("\t- %v:\n", key)
		innerVal := reflect.ValueOf(val)
		for i := 0; i < innerVal.Len(); i++ {
			color("\t\t- %v\n", innerVal.Index(i).Interface())
		}
	default:
		color("\t- %v : %v\n", key, val)
	}
}
