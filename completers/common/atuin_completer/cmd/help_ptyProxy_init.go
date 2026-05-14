package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ptyProxy_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell code to initialize atuin pty-proxy on shell startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ptyProxy_initCmd).Standalone()

	help_ptyProxyCmd.AddCommand(help_ptyProxy_initCmd)
}
