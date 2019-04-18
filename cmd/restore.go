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
	"encoding/json"
	"fmt"
	"github.com/jmg-duarte/statuspage/internal"
	"github.com/jmg-duarte/statuspage/internal/format"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var restorePath string
		if len(args) > 0 {
			restorePath = args[0]
		} else {
			restorePath = viper.GetString("backup_location")
		}

		if restorePath == "" {
			log.Fatal(fmt.Errorf("backup location not defined"))
		}

		fBackup, err := os.OpenFile(restorePath, os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		var sHist internal.ServiceHistory
		err = format.NewReader(fBackup, format.JSON).Read(&sHist)
		if err != nil {
			log.Fatal(err)
		}
		if merge {
			for service, history := range sHist {
				for t, stat := range history {
					services[service].History.AddEntry(t, stat)
				}
			}
		} else {
			for service, history := range sHist {
				services[service].History = history
			}
		}

		b, err := json.MarshalIndent(services.GetServicesHistory(), "", "\t")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(file, b, 0644)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().BoolVarP(&merge, "merge", "m", false, "")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restoreCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
