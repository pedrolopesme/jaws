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

	"github.com/spf13/cobra"
	"github.com/dgrijalva/jwt-go"
	"os"
	"github.com/fatih/color"
	"reflect"
)

var decodeCmd = &cobra.Command{
	Use:   "decode <JWT ENCODED TOKEN>",
	Short: "decodes a JWT token and Print its content",
	Long: `Decode parse an JWT token and Print its content breaking into sections. Example:

$ jaws decode <SOME JWT TOKEN>

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

		decode(args[0])
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

// TODO extract this to somewhere else
// TODO add tests
func decode(token string) {
	claims := jwt.MapClaims{}
	parsedToken, _ := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})

	Print("HEADER", parsedToken.Header, color.Magenta)
	Print("BODY", claims, color.Cyan)
}

// Print outputs a section of a JWT token, supporting title and a color.
//
//	- title : section name
//	- output: content to be printed
//  - color: color spec from fatih/color package
// TODO refactor this crap.
func Print(title string, output map[string]interface{}, color func(format string, a ...interface{})) {
	color("\n%v:", title)
	for key, val := range output {
		switch val.(type) {
		case []interface{}:
			color("\t- %v:\n", key)
			for i := 0; i < reflect.ValueOf(val).Len(); i++ {
				strct := reflect.ValueOf(val).Index(i).Interface()
				color("\t\t- %v\n", strct)
			}
		default:
			color("\t- %v : %v\n", key, val)
		}
	}
}
