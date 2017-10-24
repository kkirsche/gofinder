// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"sync"

	"github.com/kkirsche/gofinder/gofinder"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ua      string
	te      bool
	le      bool
	ce      bool
	se      bool
	quiet   bool
	verbose bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gofinder",
	Short: "Find portions of a page such as Title, Links or Comments",
	Long: `gofinder allows for you to quickly pull out details of a page. This
includes the page title, link URL's, inline and external scripts, as well as
comments. The goal of this is to simplify security auditing of sites by being
able to ignore some of the standard boilerplate nodes and focus on the ones
that are commonly of interest.`,
	Args: cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if quiet && verbose {
			logrus.Errorln("Cannot have Verbose and Quiet enabled together")
			return
		} else if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		config := gofinder.NewConfig(ua, quiet, te, le, ce, se)
		client := gofinder.NewClient(config)

		var wg sync.WaitGroup

		for _, arg := range args {
			wg.Add(1)
			go func(arg string) {
				defer wg.Done()
				resp, err := client.Get(arg)
				if err != nil {
					return
				}

				client.Find(resp)
			}(arg)
		}

		wg.Wait()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.SetUsageTemplate(`Usage:{{if .Runnable}}
  {{.UseLine}} URLs...{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)

	RootCmd.Flags().BoolVarP(&te, "title", "t", false, "Find title attributes")
	RootCmd.Flags().BoolVarP(&le, "links", "l", false, "Find link attributes")
	RootCmd.Flags().BoolVarP(&ce, "comments", "c", false, "Find comment attributes")
	RootCmd.Flags().BoolVarP(&se, "scripts", "s", false, "Find inline and external script attributes")
	RootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Only say when you found something")
	RootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	RootCmd.Flags().StringVarP(&ua, "user-agent", "u", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36", "Set the user agent you wish to use")
}
