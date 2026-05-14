package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ptyProxy_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell code to initialize atuin pty-proxy on shell startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ptyProxy_initCmd).Standalone()

	ptyProxy_initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ptyProxyCmd.AddCommand(ptyProxy_initCmd)

	carapace.Gen(ptyProxy_initCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "nu"),
	)
}
