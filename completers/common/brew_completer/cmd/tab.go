package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tabCmd = &cobra.Command{
	Use:     "tab",
	Short:   "Edit tab information for installed formulae or casks",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tabCmd).Standalone()

	tabCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	tabCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	tabCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	tabCmd.Flags().BoolP("help", "h", false, "Show this message.")
	tabCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	tabCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(tabCmd)

	carapace.Gen(tabCmd).PositionalAnyCompletion(
		action.ActionList(tabCmd).FilterArgs(),
	)
}
