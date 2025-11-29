package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all secrets",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_listCmd).Standalone()

	secrets_listCmd.Flags().StringP("format", "f", "", "Display the key values of each secret in the specified format, one of: table | dotenv | json.")
	secrets_listCmd.Flags().BoolP("show", "s", false, "Display secret values in plaintext")
	secretsCmd.AddCommand(secrets_listCmd)

	carapace.Gen(secrets_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("table", "dotenv", "json"),
	})
}
