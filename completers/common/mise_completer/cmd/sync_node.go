package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sync_nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Symlink node versions installed by nvm, nodenv, or Homebrew into mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sync_nodeCmd).Standalone()

	sync_nodeCmd.Flags().Bool("brew", false, "Get tool versions from Homebrew")
	sync_nodeCmd.Flags().BoolP("help", "h", false, "Print help")
	sync_nodeCmd.Flags().Bool("nodenv", false, "Get tool versions from nodenv")
	sync_nodeCmd.Flags().Bool("nvm", false, "Get tool versions from nvm")
	syncCmd.AddCommand(sync_nodeCmd)
}
