package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quickstartCmd = &cobra.Command{
	Use:   "quickstart",
	Aliases: []string{"q"},
	Short: "Command line script to quickly create a tox config file for a Python project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quickstartCmd).Standalone()

	add_common_flags(quickstartCmd)

	// NOTE: this command has a single optional positional argument:
	// `root - folder to create the tox.ini file (default: {cwd})`
	// This is arbitrary, so no completion.

	rootCmd.AddCommand(quickstartCmd)
}
