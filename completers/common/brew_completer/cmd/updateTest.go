package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateTestCmd = &cobra.Command{
	Use:     "update-test",
	Short:   "Run a test of `brew update` with a new repository clone",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateTestCmd).Standalone()

	updateTestCmd.Flags().Bool("before", false, "Use the commit at the specified <date> as the start commit.")
	updateTestCmd.Flags().Bool("commit", false, "Use the specified <commit> as the start commit.")
	updateTestCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updateTestCmd.Flags().Bool("help", false, "Show this message.")
	updateTestCmd.Flags().Bool("keep-tmp", false, "Retain the temporary directory containing the new repository clone.")
	updateTestCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateTestCmd.Flags().Bool("to-tag", false, "Set `HOMEBREW_UPDATE_TO_TAG` to test updating between tags.")
	updateTestCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateTestCmd)
}
