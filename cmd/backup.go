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
	"fmt"
	"github.com/jmg-duarte/statuspage/internal/format"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup your local storage to another file",

	Run: func(cmd *cobra.Command, args []string) {
		var backupPath string
		if len(args) > 0 {
			backupPath = args[0]
		} else {
			backupPath = viper.GetString("backup_location")
		}

		if backupPath == "" {
			log.Fatal(fmt.Errorf("backup location not defined"))
		}

		if _, err := os.Stat(backupPath); os.IsNotExist(err) {
			// If the file doesn't exist, create it
			_, err = os.OpenFile(backupPath, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}

		fBackup, err := os.OpenFile(backupPath, os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		switch strings.ToLower(fileFormat) {
		case "txt":
			format.NewWriter(fBackup, format.TXT)
		case "csv":
			w := format.NewWriter(fBackup, format.CSV)
			for serviceId, history := range services.GetServicesHistory() {
				w.Write([][]string{{fmt.Sprintf("[%s]", serviceId)}})
				w.Write(history.CSV())
			}
		default:
			format.NewWriter(fBackup, format.JSON).Write(services.GetServicesHistory())
		}
	},
	Args: cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.Flags().StringVarP(&fileFormat, "fileFormat", "f", "", "")
}
