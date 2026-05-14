package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ptyProxyCmd = &cobra.Command{
	Use:   "pty-proxy",
	Short: "PTY proxy for atuin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ptyProxyCmd).Standalone()

	helpCmd.AddCommand(help_ptyProxyCmd)
}
