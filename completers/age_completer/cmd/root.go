package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "age",
	Short: "simple, modern, and secure file encryption",
	Long:  "https://github.com/FiloSottile/age",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("armor", "a", false, "Encrypt to a PEM encoded format.")
	rootCmd.Flags().BoolP("decrypt", "d", false, "Decrypt the input to the output.")
	rootCmd.Flags().BoolP("encrypt", "e", false, "Encrypt the input to the output. Default if omitted.")
	rootCmd.Flags().StringArrayP("identity", "i", nil, "Use the identity file at PATH. Can be repeated.")
	rootCmd.Flags().StringP("output", "o", "", "Write the result to the file at path OUTPUT.")
	rootCmd.Flags().BoolP("passphrase", "p", false, "Encrypt with a passphrase.")
	rootCmd.Flags().StringArrayP("recipient", "r", nil, "Encrypt to the specified RECIPIENT. Can be repeated.")
	rootCmd.Flags().StringArrayP("recipients-file", "R", nil, "Encrypt to recipients listed at PATH. Can be repeated.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"identity": carapace.Batch(
			ssh.ActionPrivateKeys(),
			carapace.ActionFiles(),
		).ToA(),
		"output": carapace.ActionFiles(),
		"recipients-file": carapace.Batch(
			ssh.ActionPublicKeys(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
