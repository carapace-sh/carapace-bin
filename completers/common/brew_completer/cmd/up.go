package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Fetch the newest version of Homebrew and all formulae from GitHub using `git`(1) and perform any necessary migrations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().Bool("auto-update", false, "Run on auto-updates (e.g. before `brew install`). Skips some slower steps.")
	upCmd.Flags().Bool("debug", false, "Display a trace of all shell commands as they are executed.")
	upCmd.Flags().Bool("force", false, "Always do a slower, full update check (even if unnecessary).")
	upCmd.Flags().Bool("help", false, "Show this message.")
	upCmd.Flags().Bool("merge", false, "Use `git merge` to apply updates (rather than `git rebase`).")
	upCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	upCmd.Flags().Bool("verbose", false, "Print the directories checked and `git` operations performed.")
	rootCmd.AddCommand(upCmd)
}
