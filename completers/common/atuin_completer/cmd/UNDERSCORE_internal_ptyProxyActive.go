package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var __internal_ptyProxyActiveCmd = &cobra.Command{
	Use:   "pty-proxy-active",
	Short: "Check whether the current terminal belongs to a live PTY proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(__internal_ptyProxyActiveCmd).Standalone()

	__internalCmd.AddCommand(__internal_ptyProxyActiveCmd)
}
