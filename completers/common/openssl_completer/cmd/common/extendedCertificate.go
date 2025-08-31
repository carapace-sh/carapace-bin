package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddExtendedCertificateOptions(cmd *cobra.Command) {
	cmd.Flags().StringS("xcert", "xcert", "", "cert for Extended certificates")
	cmd.Flags().StringS("xcertform", "xcertform", "", "format of Extended certificate (PEM/DER/P12); has no effect")
	cmd.Flags().StringS("xchain", "xchain", "", "chain for Extended certificates")
	cmd.Flags().BoolS("xchain_build", "xchain_build", false, "build certificate chain for the extended certificates")
	cmd.Flags().StringS("xkey", "xkey", "", "key for Extended certificates")
	cmd.Flags().StringS("xkeyform", "xkeyform", "", "format of Extended certificate's key (DER/PEM/P12); has no effect")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"xcert":     carapace.ActionFiles(),
		"xcertform": carapace.ActionValues("DER", "PEM", "P12"),
		"xchain":    carapace.ActionFiles(),
		"xkey":      carapace.ActionFiles(),
		"xkeyform":  carapace.ActionValues("DER", "PEM", "P12"),
	})
}
