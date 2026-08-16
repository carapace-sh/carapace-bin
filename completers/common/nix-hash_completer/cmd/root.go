package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nix-hash",
	Short: "compute the cryptographic hash of a path",
	Long:  "https://nixos.org/manual/nix/stable/command-ref/nix-hash.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("base16", false, "Print hash in hexadecimal representation")
	rootCmd.Flags().Bool("base32", false, "Print hash in base-32 representation")
	rootCmd.Flags().Bool("base64", false, "Print hash in base-64 representation")
	rootCmd.Flags().Bool("flat", false, "Hash each regular file directly")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("sri", false, "Print hash in SRI format")
	rootCmd.Flags().Bool("to-base16", false, "Convert base-32 hash to hexadecimal")
	rootCmd.Flags().Bool("to-base32", false, "Convert hexadecimal hash to base-32")
	rootCmd.Flags().Bool("to-base64", false, "Convert hexadecimal hash to base-64")
	rootCmd.Flags().Bool("to-sri", false, "Convert hexadecimal hash to SRI format")
	rootCmd.Flags().Bool("truncate", false, "Truncate hashes longer than 160 bits")
	rootCmd.Flags().String("type", "", "Hash algorithm to use")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512", "blake3"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}