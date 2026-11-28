package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureGenerateKeyDigestCmd = &cobra.Command{
	Use:   "secure-generate-key-digest",
	Short: "Generate key digest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureGenerateKeyDigestCmd).Standalone()

	secureGenerateKeyDigestCmd.Flags().StringP("keyfile", "k", "", "Public key file for digest generation")
	secureGenerateKeyDigestCmd.Flags().StringP("output", "o", "", "Output file for key digest")

	rootCmd.AddCommand(secureGenerateKeyDigestCmd)

	carapace.Gen(secureGenerateKeyDigestCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})
}
