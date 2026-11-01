package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Print the stored authentication token for a given host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_tokenCmd).Standalone()

	help_authCmd.AddCommand(help_auth_tokenCmd)
}
