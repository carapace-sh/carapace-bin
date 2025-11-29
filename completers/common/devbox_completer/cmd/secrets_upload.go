package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_uploadCmd = &cobra.Command{
	Use:   "upload <file1> [<fileN>]...",
	Short: "Upload variables defined in one or more .env files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_uploadCmd).Standalone()

	secrets_uploadCmd.Flags().StringP("format", "f", "", "File format: dotenv or json")
	secretsCmd.AddCommand(secrets_uploadCmd)

	carapace.Gen(secrets_uploadCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("dotenv", "json"),
	})

	// TODO positional completion
}
