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

	updateResetCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updateResetCmd.Flags().Bool("help", false, "Show this message.")
	updateResetCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateResetCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateResetCmd)
}
