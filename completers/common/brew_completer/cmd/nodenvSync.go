package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nodenvSyncCmd = &cobra.Command{
	Use:     "nodenv-sync",
	Short:   "Create symlinks for Homebrew's installed NodeJS versions in `~/.nodenv/versions`",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nodenvSyncCmd).Standalone()

	nodenvSyncCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	nodenvSyncCmd.Flags().BoolP("help", "h", false, "Show this message.")
	nodenvSyncCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	nodenvSyncCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(nodenvSyncCmd)
}
