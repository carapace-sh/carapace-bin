package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ptyProxyCmd = &cobra.Command{
	Use:     "pty-proxy",
	Short:   "PTY proxy for atuin",
	Aliases: []string{"hex"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ptyProxyCmd).Standalone()

	ptyProxyCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(ptyProxyCmd)
}
