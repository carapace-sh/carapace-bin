package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var skeyutlCmd = &cobra.Command{
	Use:     "skeyutl",
	Short:   "opaque symmetric keys routines",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skeyutlCmd).Standalone()

	skeyutlCmd.Flags().StringS("cipher", "cipher", "", "The cipher to generate key for")
	skeyutlCmd.Flags().BoolS("genkey", "genkey", false, "Generate an opaque symmetric key")
	skeyutlCmd.Flags().StringS("skeymgmt", "skeymgmt", "", "Symmetric key management name for opaque keys handling")
	skeyutlCmd.Flags().StringSliceS("skeyopt", "skeyopt", nil, "Key options as opt:value for opaque keys handling")
	common.AddProviderFlags(skeyutlCmd)
	rootCmd.AddCommand(skeyutlCmd)
}
