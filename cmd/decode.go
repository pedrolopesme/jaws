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

	"os"

	"github.com/fatih/color"
	"github.com/pedrolopeme/jaws/internal/decode"
	"github.com/pedrolopeme/jaws/internal/model"
	"github.com/pedrolopeme/jaws/internal/utils"
	"github.com/spf13/cobra"
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

		token := decode.Decode(jwt, key)
		print(token)
	},
}

func print(token *model.Token) {
	utils.Print("HEADER", token.Header, color.Magenta)
	utils.Print("BODY", token.Claims, color.Cyan)
	fmt.Println()
}

// Print the signature info
//
//   - valid : whether the signature is valid or not
//   - key: signature key
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

func init() {
	decodeCmd.Flags().StringP("key", "k", "", "Key to validate signature")
	rootCmd.AddCommand(decodeCmd)
}
