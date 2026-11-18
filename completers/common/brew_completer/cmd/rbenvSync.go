package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbenvSyncCmd = &cobra.Command{
	Use:     "rbenv-sync",
	Short:   "Create symlinks for Homebrew's installed Ruby versions in `~/.rbenv/versions`",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbenvSyncCmd).Standalone()

	rbenvSyncCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	rbenvSyncCmd.Flags().BoolP("help", "h", false, "Show this message.")
	rbenvSyncCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	rbenvSyncCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(rbenvSyncCmd)
}
