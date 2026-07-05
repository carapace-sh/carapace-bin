package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var addCertificatesCmd = &cobra.Command{
	Use:   "add-certificates",
	Short: "Add certificates to a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCertificatesCmd).Standalone()
	addCertificatesCmd.Flags().StringP("keychain", "k", "", "Use keychain rather than the default keychain")
	rootCmd.AddCommand(addCertificatesCmd)
	carapace.Gen(addCertificatesCmd).PositionalCompletion(carapace.ActionFiles())
	carapace.Gen(addCertificatesCmd).FlagCompletion(carapace.ActionMap{
		"keychain": security.ActionKeychains(),
	})
}
