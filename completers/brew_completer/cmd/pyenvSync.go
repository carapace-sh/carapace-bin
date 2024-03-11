package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pyenvSyncCmd = &cobra.Command{
	Use:     "pyenv-sync",
	Short:   "Create symlinks for Homebrew's installed Python versions in `~/.pyenv/versions`",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pyenvSyncCmd).Standalone()

	pyenvSyncCmd.Flags().Bool("debug", false, "Display any debugging information.")
	pyenvSyncCmd.Flags().Bool("help", false, "Show this message.")
	pyenvSyncCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	pyenvSyncCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(pyenvSyncCmd)
}
