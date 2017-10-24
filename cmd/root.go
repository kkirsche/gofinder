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

	"github.com/kkirsche/gofinder/gofinder"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	title   bool
	link    bool
	comment bool
	script  bool
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
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		config := gofinder.NewConfig(title, link, comment, script)
		client := gofinder.NewClient(config)

		for _, arg := range args {
			resp, err := client.Get(arg)
			if err != nil {
				continue
			}

			client.Find(resp)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.WithError(err).Errorln("Failed to start command")
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

	RootCmd.Flags().BoolVarP(&title, "title", "t", false, "Find title attributes")
	RootCmd.Flags().BoolVarP(&link, "link", "l", false, "Find link attributes")
	RootCmd.Flags().BoolVarP(&comment, "comment", "c", false, "Find comment attributes")
	RootCmd.Flags().BoolVarP(&script, "script", "s", false, "Find inline and external script attributes")
	RootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
}
