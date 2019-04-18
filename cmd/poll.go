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
	"github.com/spf13/cobra"
	"log"
	"os"
)

// pollCmd represents the poll command
var pollCmd = &cobra.Command{
	Use:   "poll",
	Short: "Retrieves the status from of all configured services",
	Run: func(cmd *cobra.Command, args []string) {
		localStorage, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		servicesMap := internal.ValidateFilterFlags(only, exclude, services)
		err = servicesMap.PollServices(brief, localStorage)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pollCmd)

	pollCmd.Flags().StringVarP(&only, "only", "o", "", "Only poll the given services (separated by commas), takes precedence over --exclude")
	pollCmd.Flags().StringVarP(&exclude, "exclude", "e", "", "Exclude the given services (separated by commas) from being polled")
	pollCmd.Flags().BoolVarP(&brief, "brief", "b", false, "Shows only a brief overview of the results")
}
