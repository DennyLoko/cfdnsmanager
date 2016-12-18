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
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list-zones",
	Short: "List available zones",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Searching zones...")

		zones, err := cfAPI.ListZones()
		if err != nil {
			printError(err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "ID\tZONE\tPLAN")

		for _, zone := range zones {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s", zone.ID, zone.Name, zone.Plan.Name))
		}

		w.Flush()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
