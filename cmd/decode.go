// Copyright © 2018 Pedro Mendes <pedrolopesme@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"reflect"
)

var decodeCmd = &cobra.Command{
	Use:   "decode <JWT ENCODED TOKEN>",
	Short: "decodes a JWT token and Print its content",
	Long: `DecodeAndPrint parse an JWT token and Print its content breaking into sections. Example:

$ jaws decode <SOME JWT TOKEN> -k <SOME SIGNING KEY>

Header:
	- key 1: value
	- key 2: value

Body:
	- key 1: value
	- key 2: value
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must provide a JWT encoded token")
			os.Exit(126)
		}

		jwt := args[0]
		var key string

		if cmd.Flag("key") != nil {
			key, _ = cmd.Flags().GetString("key")
		}

		DecodeAndPrint(jwt, key)
	},
}

func init() {
	decodeCmd.Flags().StringP("key", "k", "", "Key to validate signature")
	rootCmd.AddCommand(decodeCmd)
}

// DecodeAndPrint prints the content of an encoded JWT .
//
//  TODO add tests
//	- token : encoded jwt
//	- key: secret key
func DecodeAndPrint(token string, key string) {
	claims := jwt.MapClaims{}
	parsedToken, _ := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	Print("HEADER", parsedToken.Header, color.Magenta)
	Print("BODY", claims, color.Cyan)
	PrintSignature(parsedToken.Valid, key)
	fmt.Println()
}

// Print outputs a section of a JWT token, supporting title and a color.
//
//	- title : section name
//	- output: content to be printed
//  - color: color spec from fatih/color package
func Print(title string, output map[string]interface{}, color func(format string, a ...interface{})) {
	color("\n%v:", title)
	for key, val := range output {
		PrintLine(key, val, color)
	}
}

// PrintLine: Output lines, formatting according to its content.
//
//	- key : claim name
//	- val: claim content
//  - color: color spec from fatih/color package
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


// Print the signature info
//
//	- valid : whether the signature is valid or not
//	- key: signature key
func PrintSignature(valid bool, key string) {
	var outputColor func(format string, a ...interface{})

	if valid {
		outputColor = color.Blue
	} else {
		outputColor = color.Red
	}

	outputColor("\nSIGNATURE:")
	outputColor("\t- VALID: %v", valid)

	if key == "" {
		outputColor("\t- REASON: No signing key provided")
		return
	}
}
