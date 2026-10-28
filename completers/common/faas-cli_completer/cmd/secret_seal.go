package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_sealCmd = &cobra.Command{
	Use:   "seal",
	Short: "Seal build secrets into an encrypted file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_sealCmd).Standalone()

	secret_sealCmd.Flags().StringArray("from-file", nil, "Secret from file in key=path format (can be repeated)")
	secret_sealCmd.Flags().StringArray("from-literal", nil, "Literal secret in key=value format (can be repeated)")
	secret_sealCmd.Flags().StringP("output", "o", "com.openfaas.secrets", "Output file path")
	secretCmd.AddCommand(secret_sealCmd)

	carapace.Gen(secret_sealCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(secret_sealCmd).PositionalCompletion(
		carapace.ActionFiles(".pub"),
	)
}
