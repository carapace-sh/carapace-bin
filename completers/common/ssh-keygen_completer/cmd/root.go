package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ssh-keygen_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/ssh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssh-keygen",
	Short: "OpenSSH authentication key utility",
	Long:  "https://linux.die.net/man/1/ssh-keygen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Generate missing host keys for each of the key types")
	rootCmd.Flags().BoolS("B", "B", false, "Show the bubblebabble digest of specified private or public key file")
	rootCmd.Flags().StringS("C", "C", "", "Provides a new comment")
	rootCmd.Flags().StringS("D", "D", "", "Download the public keys provided by the PKCS#11 shared library pkcs11")
	rootCmd.Flags().StringS("E", "E", "", "Specifies the hash algorithm used when displaying key fingerprints")
	rootCmd.Flags().StringS("F", "F", "", "Search for the specified hostname in a known_hosts file")
	rootCmd.Flags().BoolS("H", "H", false, "Hash a known_hosts file")
	rootCmd.Flags().StringS("I", "I", "", "Specify the key identity when signing a public key")
	rootCmd.Flags().BoolS("K", "K", false, "Download resident keys from a FIDO authenticator")
	rootCmd.Flags().BoolS("L", "L", false, "Prints the contents of one or more certificates")
	rootCmd.Flags().StringS("M", "M", "", "Moduli")
	rootCmd.Flags().StringS("N", "N", "", "Provides the new passphrase")
	rootCmd.Flags().StringSliceS("O", "O", nil, "Specify a key/value option")
	rootCmd.Flags().StringS("P", "P", "", "Provides the (old) passphrase")
	rootCmd.Flags().BoolS("Q", "Q", false, "Test whether keys have been revoked in a KRL")
	rootCmd.Flags().StringS("R", "R", "", "Removes all keys belonging to the specified hostname from a known_hosts file")
	rootCmd.Flags().BoolS("U", "U", false, "Indicates that a CA key resides in a ssh-agent")
	rootCmd.Flags().StringS("V", "V", "", "Specify a validity interval when signing a certificate")
	rootCmd.Flags().StringS("Y", "Y", "", "Operation")
	rootCmd.Flags().StringS("Z", "Z", "", "Specifies the cipher to use for encryption when writing an OpenSSH-format private key file")
	rootCmd.Flags().StringS("a", "a", "", "Number of KDF rounds used when saving a private key")
	rootCmd.Flags().StringS("b", "b", "", "Specifies the number of bits in the key to create")
	rootCmd.Flags().BoolS("c", "c", false, "Requests changing the comment in the private and public key files")
	rootCmd.Flags().BoolS("e", "e", false, "This option will read a private or public OpenSSH key file and print to stdout")
	rootCmd.Flags().StringS("f", "f", "", "Specifies the filename of the key file")
	rootCmd.Flags().BoolS("g", "g", false, "Use generic DNS format when printing fingerprint resource records")
	rootCmd.Flags().BoolS("h", "h", false, "When signing a key, create a host certificate instead of a user certificate")
	rootCmd.Flags().BoolS("i", "i", false, "Read an unencrypted key file and print an OpenSSH compatible key to stdout")
	rootCmd.Flags().BoolS("k", "k", false, "Generate a KRL file")
	rootCmd.Flags().BoolS("l", "l", false, "Show fingerprint of specified public key file")
	rootCmd.Flags().StringS("m", "m", "", "Specify a key format for key generation")
	rootCmd.Flags().StringS("n", "n", "", "Specify one or more principals to be included in a certificate")
	rootCmd.Flags().BoolS("p", "p", false, "Requests changing the passphrase of a private key file")
	rootCmd.Flags().BoolS("q", "q", false, "Silence ssh-keygen")
	rootCmd.Flags().StringS("r", "r", "", "Print the SSHFP fingerprint resource record named hostname for the specified public key file")
	rootCmd.Flags().StringS("s", "s", "", "Certify (sign) a public key using the specified CA key")
	rootCmd.Flags().StringS("t", "t", "", "Specifies the type of key to create")
	rootCmd.Flags().BoolS("u", "u", false, "Update a KRL")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")
	rootCmd.Flags().StringS("w", "w", "", "Specifies a path to a library that will be used when creating FIDO authenticator-hosted keys")
	rootCmd.Flags().BoolS("y", "y", false, "This option will read a private OpenSSH format file and print an OpenSSH public key to stdout")
	rootCmd.Flags().StringS("z", "z", "", "Specifies a serial number to be embedded in the certificate")

	// TODO finish flag completion (especially options)
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"E": carapace.ActionValues("md5", "sha256"),
		"F": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().NoSpace()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"M": carapace.ActionValuesDescribed(
			"generate", "Generate candidate Diffie-Hellman Group Exchange (DH-GEX) parameters",
			"screen", "Screen candidate parameters for Diffie-Hellman Group Exchange",
		),
		"O": action.ActionOptions(),
		"R": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts().NoSpace()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
		"Y": carapace.ActionValuesDescribed(
			"find-principals", "Find the principal(s) associated with the public key of a signature",
			"match-principals", "Find principal matching the principal name provided using the -I flag",
			"check-novalidate", "Checks that a signature generated using ssh-keygen -Y sign has a valid structure",
			"sign", "Cryptographically sign a file or some data using a SSH key",
			"verify", "Request to verify a signature generated using ssh-keygen -Y sign",
		),
		"Z": ssh.ActionCiphers(),
		"b": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch rootCmd.Flag("t").Value.String() {
			case "rsa":
				return carapace.ActionValues("1024", "2048", "3072")
			case "dsa":
				return carapace.ActionValues("1024")
			case "ecdsa":
				return carapace.ActionValues("256", "384", "521")
			default:
				return carapace.ActionValues()
			}
		}),
		"f": carapace.ActionFiles(),
		"m": carapace.ActionValuesDescribed(
			"RFC4716", "RFC 4716/SSH2 public or private key",
			"PKCS8", "PKCS8 public or private key",
			"PEM", "PEM public key",
		),
		"n": carapace.Batch(
			os.ActionUsers(),
			net.ActionHosts(),
		).ToA().UniqueList(","),
		"r": net.ActionHosts(),
		"s": carapace.ActionFiles(),
		"t": carapace.ActionValues("dsa", "ecdsa", "ecdsa-sk", "ed25519", "ed25519-sk", "rsa"),
		"w": carapace.ActionFiles(),
	})
}
