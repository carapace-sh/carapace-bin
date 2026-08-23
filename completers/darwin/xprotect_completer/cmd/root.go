package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xprotect",
	Short: "utility for interacting with XProtect",
	Long:  "https://keith.github.io/xcode-manpages/xprotect.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"update", "Perform an update of XProtect assets",
			"check", "Print the currently online available update version",
			"version", "Print the version of the currently installed XProtect assets",
			"logs", "Display XProtect logs",
			"status", "Print the current status of XProtect",
			"help", "Print help for a particular subcommand",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if c.Args[0] == "update" {
				return carapace.ActionValues("--prerelease")
			}
			if c.Args[0] == "check" || c.Args[0] == "version" || c.Args[0] == "status" {
				return carapace.ActionValues("--json")
			}
			return carapace.ActionValues()
		}),
	)
}