package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sync_rubyCmd = &cobra.Command{
	Use:   "ruby",
	Short: "Symlink ruby versions installed by Homebrew into mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sync_rubyCmd).Standalone()

	sync_rubyCmd.Flags().Bool("brew", false, "Get tool versions from Homebrew")
	sync_rubyCmd.Flags().BoolP("help", "h", false, "Print help")
	sync_rubyCmd.MarkFlagRequired("brew")
	syncCmd.AddCommand(sync_rubyCmd)
}
