package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Fetch the newest version of Homebrew and all formulae from GitHub using `git`(1) and perform any necessary migrations",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("auto-update", false, "Run on auto-updates (e.g. before `brew install`). Skips some slower steps.")
	updateCmd.Flags().Bool("debug", false, "Display a trace of all shell commands as they are executed.")
	updateCmd.Flags().Bool("force", false, "Always do a slower, full update check (even if unnecessary).")
	updateCmd.Flags().Bool("help", false, "Show this message.")
	updateCmd.Flags().Bool("merge", false, "Use `git merge` to apply updates (rather than `git rebase`).")
	updateCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateCmd.Flags().Bool("verbose", false, "Print the directories checked and `git` operations performed.")
	rootCmd.AddCommand(updateCmd)
}
