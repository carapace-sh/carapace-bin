package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var encCmd = &cobra.Command{
	Use:     "enc",
	Short:   "Encryption, decryption, and encoding",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(encCmd).Standalone()

	encCmd.Flags().BoolS("A", "A", false, "Used with -[base64|a] to specify base64 buffer as a single line")
	encCmd.Flags().StringS("K", "K", "", "Raw key, in hex")
	encCmd.Flags().BoolS("P", "P", false, "Print the iv/key and exit")
	encCmd.Flags().StringS("S", "S", "", "Salt, in hex")
	encCmd.Flags().BoolS("a", "a", false, "Base64 encode/decode, depending on encryption flag")
	encCmd.Flags().BoolS("base64", "base64", false, "Same as option -a")
	encCmd.Flags().StringS("bufsize", "bufsize", "", "Buffer size")
	encCmd.Flags().BoolS("ciphers", "ciphers", false, "Alias for -list")
	encCmd.Flags().BoolS("d", "d", false, "Decrypt")
	encCmd.Flags().BoolS("debug", "debug", false, "Print debug info")
	encCmd.Flags().BoolS("e", "e", false, "Encrypt")
	encCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	encCmd.Flags().StringS("in", "in", "", "Input file")
	encCmd.Flags().StringS("iter", "iter", "", "Specify the iteration count and force the use of PBKDF2 Default: 10000")
	encCmd.Flags().StringS("iv", "iv", "", "IV in hex")
	encCmd.Flags().StringS("k", "k", "", "Passphrase")
	encCmd.Flags().StringS("kfile", "kfile", "", "Read passphrase from file")
	encCmd.Flags().BoolS("list", "list", false, "List ciphers")
	encCmd.Flags().StringS("md", "md", "", "Use specified digest to create a key from the passphrase")
	encCmd.Flags().BoolS("none", "none", false, "Don't encrypt")
	encCmd.Flags().BoolS("nopad", "nopad", false, "Disable standard block padding")
	encCmd.Flags().BoolS("nosalt", "nosalt", false, "Do not use salt in the KDF")
	encCmd.Flags().StringS("out", "out", "", "Output file")
	encCmd.Flags().BoolS("p", "p", false, "Print the iv/key")
	encCmd.Flags().StringS("pass", "pass", "", "Passphrase source")
	encCmd.Flags().BoolS("pbkdf2", "pbkdf2", false, "Use password-based key derivation function 2 (PBKDF2) Use -iter to change the iteration count from 10000")
	encCmd.Flags().BoolS("salt", "salt", false, "Use salt in the KDF (default)")
	encCmd.Flags().StringS("saltlen", "saltlen", "", "Specify the PBKDF2 salt length (in bytes) Default: 16")
	encCmd.Flags().StringS("skeymgmt", "skeymgmt", "", "Symmetric key management name for opaque symmetric key handling")
	encCmd.Flags().StringSliceS("skeyopt", "skeyopt", nil, "Key options as opt:value for opaque symmetric key handling")
	encCmd.Flags().BoolS("v", "v", false, "Verbose output")
	encCmd.Flags().BoolS("z", "z", false, "Compress or decompress encrypted data using zlib")
	common.AddProviderFlags(encCmd)
	common.AddRandomStateFlags(encCmd)
	rootCmd.AddCommand(encCmd)

	carapace.Gen(encCmd).FlagCompletion(carapace.ActionMap{
		"engine": action.ActionEngines(),
		"in":     carapace.ActionFiles(),
		"kfile":  carapace.ActionFiles(),
		"md":     action.ActionDigestAlgorithms(encCmd),
		"out":    carapace.ActionFiles(),
	})
}
