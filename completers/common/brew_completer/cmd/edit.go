package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Open a <formula>, <cask> or <tap> in the editor set by `EDITOR` or `HOMEBREW_EDITOR`, or open the Homebrew repository for editing if no argument is provided",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	editCmd.Flags().Bool("debug", false, "Display any debugging information.")
	editCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	editCmd.Flags().Bool("help", false, "Show this message.")
	editCmd.Flags().Bool("print-path", false, "Print the file path to be edited, without opening an editor.")
	editCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	editCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalAnyCompletion(
		action.ActionList(editCmd),
	)
}
