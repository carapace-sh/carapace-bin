package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "list algorithms and features",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolS("1", "1", false, "List in one column")
	listCmd.Flags().BoolS("all-algorithms", "all-algorithms", false, "List of all algorithms")
	listCmd.Flags().BoolS("all-tls-groups", "all-tls-groups", false, "List implemented TLS key exchange 'groups' and all aliases")
	listCmd.Flags().BoolS("asymcipher-algorithms", "asymcipher-algorithms", false, "List of asymmetric cipher algorithms")
	listCmd.Flags().BoolS("cipher-algorithms", "cipher-algorithms", false, "List of symmetric cipher algorithms")
	listCmd.Flags().BoolS("cipher-commands", "cipher-commands", false, "List of cipher commands (deprecated)")
	listCmd.Flags().BoolS("commands", "commands", false, "List of standard commands")
	listCmd.Flags().BoolS("decoders", "decoders", false, "List of decoding methods")
	listCmd.Flags().BoolS("digest-algorithms", "digest-algorithms", false, "List of message digest algorithms")
	listCmd.Flags().BoolS("digest-commands", "digest-commands", false, "List of message digest commands (deprecated)")
	listCmd.Flags().BoolS("disabled", "disabled", false, "List of disabled features")
	listCmd.Flags().BoolS("encoders", "encoders", false, "List of encoding methods")
	listCmd.Flags().BoolS("kdf-algorithms", "kdf-algorithms", false, "List of key derivation and pseudo random function algorithms")
	listCmd.Flags().BoolS("kem-algorithms", "kem-algorithms", false, "List of key encapsulation mechanism algorithms")
	listCmd.Flags().BoolS("key-exchange-algorithms", "key-exchange-algorithms", false, "List of key exchange algorithms")
	listCmd.Flags().BoolS("key-managers", "key-managers", false, "List of key managers")
	listCmd.Flags().BoolS("mac-algorithms", "mac-algorithms", false, "List of message authentication code algorithms")
	listCmd.Flags().BoolS("objects", "objects", false, "List built in objects (OID<->name mappings)")
	listCmd.Flags().StringS("options", "options", "", "List options for specified command")
	listCmd.Flags().BoolS("providers", "providers", false, "List of provider information")
	listCmd.Flags().BoolS("public-key-algorithms", "public-key-algorithms", false, "List of public key algorithms")
	listCmd.Flags().BoolS("public-key-methods", "public-key-methods", false, "List of public key methods")
	listCmd.Flags().BoolS("random-generators", "random-generators", false, "List of random number generators")
	listCmd.Flags().BoolS("random-instances", "random-instances", false, "List the primary, public and private random number generator details")
	listCmd.Flags().StringS("select", "select", "", "Select a single algorithm")
	listCmd.Flags().BoolS("signature-algorithms", "signature-algorithms", false, "List of signature algorithms")
	listCmd.Flags().BoolS("skey-managers", "skey-managers", false, "List of symmetric key managers")
	listCmd.Flags().BoolS("standard-commands", "standard-commands", false, "List of standard commands")
	listCmd.Flags().BoolS("store-loaders", "store-loaders", false, "List of store loaders")
	listCmd.Flags().BoolS("tls-groups", "tls-groups", false, "List implemented TLS key exchange 'groups'")
	listCmd.Flags().BoolS("tls-signature-algorithms", "tls-signature-algorithms", false, "List of TLS signature algorithms")
	listCmd.Flags().BoolS("tls1_2", "tls1_2", false, "When listing 'groups', list those compatible with TLS1.2")
	listCmd.Flags().BoolS("tls1_3", "tls1_3", false, "When listing 'groups', list those compatible with TLS1.3")
	listCmd.Flags().BoolS("verbose", "verbose", false, "Verbose listing")
	common.AddProviderFlags(listCmd)
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"options": carapace.ActionCommands(rootCmd),
	})
}
