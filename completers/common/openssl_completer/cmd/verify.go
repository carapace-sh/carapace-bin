package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:     "verify",
	Short:   "X.509 Certificate Verification",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()

	verifyCmd.Flags().StringS("CAfile", "CAfile", "", "A file of trusted certificates")
	verifyCmd.Flags().StringS("CApath", "CApath", "", "A directory of files with trusted certificates")
	verifyCmd.Flags().StringS("CAstore", "CAstore", "", "URI to a store of trusted certificates")
	verifyCmd.Flags().StringS("CRLfile", "CRLfile", "", "File containing one or more CRL's (in PEM format) to load")
	verifyCmd.Flags().BoolS("crl_download", "crl_download", false, "Try downloading CRL information for certificates via their CDP entries")
	verifyCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	verifyCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default trusted certificates file")
	verifyCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load trusted certificates from the default directory")
	verifyCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load trusted certificates from the default certificates store")
	verifyCmd.Flags().BoolS("show_chain", "show_chain", false, "Display information about the certificate chain")
	verifyCmd.Flags().StringS("trusted", "trusted", "", "A file of trusted certificates")
	verifyCmd.Flags().StringS("untrusted", "untrusted", "", "A file of untrusted certificates")
	verifyCmd.Flags().BoolS("verbose", "verbose", false, "Print extra information about the operations being performed.")
	verifyCmd.Flags().StringSliceS("vfyopt", "vfyopt", nil, "Verification parameter in n:v form")
	common.AddProviderFlags(verifyCmd)
	common.AddValidationFlags(verifyCmd)
	rootCmd.AddCommand(verifyCmd)

	carapace.Gen(verifyCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":    carapace.ActionFiles(),
		"CApath":    carapace.ActionDirectories(),
		"CRLfile":   carapace.ActionFiles(),
		"trusted":   carapace.ActionFiles(),
		"untrusted": carapace.ActionFiles(),
	})

	carapace.Gen(verifyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
