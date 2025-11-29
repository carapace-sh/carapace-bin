package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_removeCmd = &cobra.Command{
	Use:     "remove <NAME1> [<NAME2>]...",
	Short:   "Remove one or more environment variables",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_removeCmd).Standalone()

	secretsCmd.AddCommand(secrets_removeCmd)

	// TODO positional completion
}
