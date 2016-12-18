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
	"strings"
	"text/tabwriter"

	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
)

var types string

var listrecordsCmd = &cobra.Command{
	Use:   "list-records [zone ID]",
	Short: "List the DNS records of a given zone",
	Long: `List the DNS records of a given zone.
Query the Zone ID with list-zones.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printNotice("You must specify the zone ID.")
			os.Exit(1)
		}

		rtypes := strings.Split(types, ",")
		for i, rtype := range rtypes {
			fmt.Println(fmt.Sprintf("Searching for %s records...", rtype))

			records, err := cfAPI.DNSRecords(args[0], cloudflare.DNSRecord{
				Type: rtype,
			})
			if err != nil {
				printError(err)
				os.Exit(1)
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', tabwriter.TabIndent)
			fmt.Fprintln(w, "ID\tTYPE\tPROXIED\tTTL\tNAME\tCONTENT")

			for _, record := range records {
				proxied := "Y"
				if !record.Proxied {
					proxied = "N"
				}

				fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t%d\t%s\t%s",
					record.ID,
					record.Type,
					proxied,
					record.TTL,
					record.Name,
					record.Content,
				))
			}

			w.Flush()

			i = i + 1
			if i < len(rtypes) {
				fmt.Print("\n")
			}
		}
	},
}

func init() {
	listrecordsCmd.Flags().StringVar(&types, "type", "A,NS,MX,CNAME", "DNS record type to query")

	RootCmd.AddCommand(listrecordsCmd)
}
