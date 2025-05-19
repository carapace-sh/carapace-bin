package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show all secrets in a list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_lsCmd).Standalone()

	secretsCmd.AddCommand(secrets_lsCmd)
}
