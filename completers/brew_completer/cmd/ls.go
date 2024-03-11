package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all installed formulae and casks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().BoolS("1", "1", false, "Force output to be one entry per line. This is the default when output is not to a terminal.")
	lsCmd.Flags().Bool("cask", false, "List only casks, or treat all named arguments as casks.")
	lsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	lsCmd.Flags().Bool("formula", false, "List only formulae, or treat all named arguments as formulae.")
	lsCmd.Flags().Bool("full-name", false, "Print formulae with fully-qualified names. Unless `--full-name`, `--versions` or `--pinned` are passed, other options (i.e. `-1`, `-l`, `-r` and `-t`) are passed to `ls`(1) which produces the actual output.")
	lsCmd.Flags().Bool("help", false, "Show this message.")
	lsCmd.Flags().BoolS("l", "l", false, "List formulae and/or casks in long format. Has no effect when a formula or cask name is passed as an argument.")
	lsCmd.Flags().Bool("multiple", false, "Only show formulae with multiple versions installed.")
	lsCmd.Flags().Bool("pinned", false, "List only pinned formulae, or only the specified (pinned) formulae if <formula> are provided. See also `pin`, `unpin`.")
	lsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	lsCmd.Flags().BoolS("r", "r", false, "Reverse the order of the formulae and/or casks sort to list the oldest entries first. Has no effect when a formula or cask name is passed as an argument.")
	lsCmd.Flags().BoolS("t", "t", false, "Sort formulae and/or casks by time modified, listing most recently modified first. Has no effect when a formula or cask name is passed as an argument.")
	lsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	lsCmd.Flags().Bool("versions", false, "Show the version number for installed formulae, or only the specified formulae if <formula> are provided.")
	rootCmd.AddCommand(lsCmd)
}
