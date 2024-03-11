package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hash_fileCmd = &cobra.Command{
	Use:   "file",
	Short: "print cryptographic hash of a regular file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_fileCmd).Standalone()

	hash_fileCmd.Flags().Bool("base16", false, "Print the hash in base-16 format")
	hash_fileCmd.Flags().Bool("base32", false, "Print the hash in base-32 (Nix-specific) format")
	hash_fileCmd.Flags().Bool("base64", false, "Print the hash in base-64 format")
	hash_fileCmd.Flags().Bool("sri", false, "Print the hash in SRI format")
	hash_fileCmd.Flags().String("type", "", "hash algorithm")
	hashCmd.AddCommand(hash_fileCmd)

	addLoggingFlags(hash_fileCmd)

	carapace.Gen(hash_fileCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(hash_fileCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
