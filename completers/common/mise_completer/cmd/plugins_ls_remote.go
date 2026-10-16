package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugins_ls_remoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List available plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_ls_remoteCmd).Standalone()

	plugins_ls_remoteCmd.Flags().Bool("only-names", false, "Only show plugin names")
	plugins_ls_remoteCmd.Flags().BoolP("urls", "u", false, "Show git urls")
	pluginsCmd.AddCommand(plugins_ls_remoteCmd)
}
