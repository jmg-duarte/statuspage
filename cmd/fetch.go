// Copyright © 2019 José Duarte <jmg.duarte@campus.fct.unl.pt>
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
	"github.com/jmg-duarte/statuspage/internal"
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Retrieves the status from of all configured services with a given interval (default interval: 5 seconds)",
	Run: func(cmd *cobra.Command, args []string) {
		var localStorage io.WriteSeeker
		localStorage, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		servicesMap := internal.ValidateFilterFlags(only, exclude, services)
		servicesMap.FetchServices(brief, interval, localStorage)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringVarP(&only, "only", "o", "", "Only fetch the given services (separated by commas), takes precedence over --exclude")
	fetchCmd.Flags().StringVarP(&exclude, "exclude", "e", "", "Exclude the given services (separated by commas) from being fetched")
	fetchCmd.Flags().BoolVarP(&brief, "brief", "b", false, "Shows only a brief overview of the results")
	fetchCmd.Flags().DurationVarP(&interval, "refresh", "r", 5*time.Second, "Specify the refresh rate (e.g. \"-r 23s\" for 23 second intervals)")

}
