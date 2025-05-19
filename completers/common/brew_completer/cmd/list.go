package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all installed formulae and casks",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolS("1", "1", false, "Force output to be one entry per line. This is the default when output is not to a terminal.")
	listCmd.Flags().Bool("cask", false, "List only casks, or treat all named arguments as casks.")
	listCmd.Flags().Bool("debug", false, "Display any debugging information.")
	listCmd.Flags().Bool("formula", false, "List only formulae, or treat all named arguments as formulae.")
	listCmd.Flags().Bool("full-name", false, "Print formulae with fully-qualified names. Unless `--full-name`, `--versions` or `--pinned` are passed, other options (i.e. `-1`, `-l`, `-r` and `-t`) are passed to `ls`(1) which produces the actual output.")
	listCmd.Flags().Bool("help", false, "Show this message.")
	listCmd.Flags().BoolS("l", "l", false, "List formulae and/or casks in long format. Has no effect when a formula or cask name is passed as an argument.")
	listCmd.Flags().Bool("multiple", false, "Only show formulae with multiple versions installed.")
	listCmd.Flags().Bool("pinned", false, "List only pinned formulae, or only the specified (pinned) formulae if <formula> are provided. See also `pin`, `unpin`.")
	listCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	listCmd.Flags().BoolS("r", "r", false, "Reverse the order of the formulae and/or casks sort to list the oldest entries first. Has no effect when a formula or cask name is passed as an argument.")
	listCmd.Flags().BoolS("t", "t", false, "Sort formulae and/or casks by time modified, listing most recently modified first. Has no effect when a formula or cask name is passed as an argument.")
	listCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	listCmd.Flags().Bool("versions", false, "Show the version number for installed formulae, or only the specified formulae if <formula> are provided.")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		action.ActionList(listCmd).FilterArgs(),
	)
}
