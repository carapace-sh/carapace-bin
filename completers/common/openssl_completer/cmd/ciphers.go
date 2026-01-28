package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var ciphersCmd = &cobra.Command{
	Use:     "ciphers",
	Short:   "Cipher Suite Description Determination",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciphersCmd).Standalone()

	ciphersCmd.Flags().BoolS("V", "V", false, "Even more verbose")
	ciphersCmd.Flags().StringS("ciphersuites", "ciphersuites", "", "Configure the TLSv1.3 ciphersuites to use")
	ciphersCmd.Flags().StringS("convert", "convert", "", "Convert standard name into OpenSSL name")
	ciphersCmd.Flags().BoolS("psk", "psk", false, "Include ciphersuites requiring PSK")
	ciphersCmd.Flags().BoolS("s", "s", false, "Only supported ciphers")
	ciphersCmd.Flags().BoolS("srp", "srp", false, "(deprecated) Include ciphersuites requiring SRP")
	ciphersCmd.Flags().BoolS("stdname", "stdname", false, "Show standard cipher names")
	ciphersCmd.Flags().BoolS("tls1", "tls1", false, "Ciphers compatible with TLS1")
	ciphersCmd.Flags().BoolS("tls1_1", "tls1_1", false, "Ciphers compatible with TLS1.1")
	ciphersCmd.Flags().BoolS("tls1_2", "tls1_2", false, "Ciphers compatible with TLS1.2")
	ciphersCmd.Flags().BoolS("tls1_3", "tls1_3", false, "Ciphers compatible with TLS1.3")
	ciphersCmd.Flags().BoolS("v", "v", false, "Verbose listing of the SSL/TLS ciphers")
	common.AddProviderFlags(ciphersCmd)
	rootCmd.AddCommand(ciphersCmd)
}
