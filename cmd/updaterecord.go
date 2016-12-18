// Copyright © 2016 Danniel Magno
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	newTTL       int
	selfIP       bool
	continueExec = false
)

var updaterecordCmd = &cobra.Command{
	Use:   "update-record [zone ID] [record ID] [content]",
	Short: "Update a DNS record",
	Long: `Update a DNS record
Query the record ID with list-records.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 && selfIP == false {
			printNotice("You must specify the zone ID, record ID and it's new content.")
			os.Exit(1)
		} else if len(args) < 2 && selfIP == true {
			printNotice("You must specify the zone ID and record ID.")
			os.Exit(1)
		}

		var zoneID, recordID, content string
		zoneID = args[0]
		recordID = args[1]

		if !selfIP {
			content = args[2]
		} else {
			var err error
			content, err = getIP()
			if err != nil {
				printError(err)
				os.Exit(1)
			}

			fmt.Println(fmt.Sprintf("Self IP: %s", content))
		}

		err := updateRecord(zoneID, recordID, content, newTTL)
		if err != nil {
			printError(err)
			os.Exit(1)
		}

		printSuccess("Done.")
	},
}

func init() {
	updaterecordCmd.Flags().IntVar(&newTTL, "ttl", 0, "Set the new TTL of the record")
	updaterecordCmd.Flags().BoolVar(&selfIP, "selfip", false, "Set the value of the record the IP of the machine. Overwrites value argument.")

	RootCmd.AddCommand(updaterecordCmd)
}
