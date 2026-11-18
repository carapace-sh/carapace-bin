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
	upCmd.Flags().BoolP("debug", "d", false, "Display a trace of all shell commands as they are executed.")
	upCmd.Flags().BoolP("force", "f", false, "Always do a slower, full update check (even if unnecessary).")
	upCmd.Flags().BoolP("help", "h", false, "Show this message.")
	upCmd.Flags().Bool("merge", false, "Use `git merge` to apply updates (rather than `git rebase`).")
	upCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	upCmd.Flags().BoolP("verbose", "v", false, "Print the directories checked and `git` operations performed.")
	rootCmd.AddCommand(upCmd)
}
