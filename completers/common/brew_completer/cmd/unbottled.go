package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unbottledCmd = &cobra.Command{
	Use:     "unbottled",
	Short:   "Show the unbottled dependents of formulae",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unbottledCmd).Standalone()

	unbottledCmd.Flags().Bool("debug", false, "Display any debugging information.")
	unbottledCmd.Flags().Bool("dependents", false, "Skip getting analytics data and sort by number of dependents instead.")
	unbottledCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks, whether installed or not, to check them. Implied if `HOMEBREW_EVAL_ALL` is set.")
	unbottledCmd.Flags().Bool("help", false, "Show this message.")
	unbottledCmd.Flags().Bool("lost", false, "Print the `homebrew/core` commits where bottles were lost in the last week.")
	unbottledCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	unbottledCmd.Flags().Bool("tag", false, "Use the specified bottle tag (e.g. `big_sur`) instead of the current OS.")
	unbottledCmd.Flags().Bool("total", false, "Print the number of unbottled and total formulae.")
	unbottledCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(unbottledCmd)
}
