package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certs_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new certificate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_addCmd).Standalone()

	certs_addCmd.Flags().String("ca", "", "CA certificate chain file")
	certs_addCmd.Flags().String("crt", "", "Certificate file")
	certs_addCmd.Flags().String("key", "", "Certificate key file")

	certsCmd.AddCommand(certs_addCmd)

	carapace.Gen(certs_addCmd).FlagCompletion(carapace.ActionMap{
		"ca":  carapace.ActionFiles(),
		"crt": carapace.ActionFiles(),
		"key": carapace.ActionFiles(),
	})
}
