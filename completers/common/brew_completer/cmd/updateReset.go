package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateResetCmd = &cobra.Command{
	Use:     "update-reset",
	Short:   "Fetch and reset Homebrew and all tap repositories to their latest `origin/HEAD`",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateResetCmd).Standalone()

	updateResetCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	updateResetCmd.Flags().BoolP("help", "h", false, "Show this message.")
	updateResetCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	updateResetCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateResetCmd)
}
