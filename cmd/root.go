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
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile  string
	services internal.Services

	only     string
	exclude  string
	brief    bool
	interval time.Duration

	fWriter io.Writer
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "statuspage",
	Short: "A CLI tool to query statuspage.io pages",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, loadConfig, loadLocalStorage)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.statuspage.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".statuspage" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".statuspage")
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.SetConfigType("json")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func loadConfig() {
	res := viper.Get("services")
	if res == nil {
		log.Fatal("\"services\" not defined in the config file")
	}

	services = make(internal.Services)
	// Convert the resulting interface{} into []interface{}
	// Then convert it into []map[string]string
	for _, i := range res.([]interface{}) {
		s, err := cast.ToStringMapStringE(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		services.Add(s)
	}
}

func loadLocalStorage() {
	file := viper.GetString("output")
	if file == "" {
		log.Fatal("output path cannot be null/empty")
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// If the file doesn't exist, create it
		fWriter, err = os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// If the file exists try and read from it
		fWriter, err = os.OpenFile(file, os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadFile(file)
		if err != nil {
			// If there was an error reading the file, exit
			log.Fatal(err)
		}
		var sHist internal.ServiceHistory
		err = json.Unmarshal(b, &sHist)
		if err != nil {
			// If there was an error parsing it, exit
			log.Fatal(err)
		}
		// No history yet
		for id, service := range services {
			service.History = sHist[id]
		}
	}
}
