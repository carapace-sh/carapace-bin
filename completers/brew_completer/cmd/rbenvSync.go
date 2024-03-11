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

	rbenvSyncCmd.Flags().Bool("debug", false, "Display any debugging information.")
	rbenvSyncCmd.Flags().Bool("help", false, "Show this message.")
	rbenvSyncCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	rbenvSyncCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(rbenvSyncCmd)
}
