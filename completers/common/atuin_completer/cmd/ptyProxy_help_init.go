package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ptyProxy_help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell code to initialize atuin pty-proxy on shell startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ptyProxy_help_initCmd).Standalone()

	ptyProxy_helpCmd.AddCommand(ptyProxy_help_initCmd)
}
