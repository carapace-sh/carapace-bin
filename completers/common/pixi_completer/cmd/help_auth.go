package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login to prefix.dev or anaconda.org servers to access private channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_authCmd).Standalone()

	helpCmd.AddCommand(help_authCmd)
}
