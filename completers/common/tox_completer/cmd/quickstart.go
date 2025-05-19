package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quickstartCmd = &cobra.Command{
	Use:     "quickstart",
	Aliases: []string{"q"},
	Short:   "Command line script to quickly create a tox config file for a Python project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quickstartCmd).Standalone()

	rootCmd.AddCommand(quickstartCmd)

	carapace.Gen(quickstartCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
