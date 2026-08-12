package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Print the stored authentication token for a given host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_tokenCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_tokenCmd)
}
