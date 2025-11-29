package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secrets_downloadCmd = &cobra.Command{
	Use:   "download <file1>",
	Short: "Download environment variables into the specified file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_downloadCmd).Standalone()

	secrets_downloadCmd.Flags().StringP("format", "f", "", "file format: dotenv or json")
	secretsCmd.AddCommand(secrets_downloadCmd)

	carapace.Gen(secrets_downloadCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("dotenv", "json"),
	})

	// TODO positional completion
}
