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

	ptyProxyCmd.Flags().Bool("debug-markers", false, "Highlight `OSC 18188735` escape sequences used to mark the start and end of command output")
	ptyProxyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ptyProxyCmd.Flags().StringSlice("shell", nil, "Path to the shell binary that atuin pty-proxy should spawn")
	rootCmd.AddCommand(ptyProxyCmd)
}
