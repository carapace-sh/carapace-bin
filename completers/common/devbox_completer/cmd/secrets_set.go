package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_setCmd = &cobra.Command{
	Use:   "set <NAME1>=<value1> [<NAME2>=<value2>]...",
	Short: "Securely store one or more environment variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_setCmd).Standalone()

	secretsCmd.AddCommand(secrets_setCmd)

	// TODO positional completion
}
