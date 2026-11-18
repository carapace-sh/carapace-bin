package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help___internal_ptyProxyActiveCmd = &cobra.Command{
	Use:   "pty-proxy-active",
	Short: "Check whether the current terminal belongs to a live PTY proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help___internal_ptyProxyActiveCmd).Standalone()

	help___internalCmd.AddCommand(help___internal_ptyProxyActiveCmd)
}
