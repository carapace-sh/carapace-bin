package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var hash_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "print cryptographic hash of the NAR serialisation of a path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_pathCmd).Standalone()

	hash_pathCmd.Flags().Bool("base16", false, "Print the hash in base-16 format")
	hash_pathCmd.Flags().Bool("base32", false, "Print the hash in base-32 (Nix-specific) format")
	hash_pathCmd.Flags().Bool("base64", false, "Print the hash in base-64 format")
	hash_pathCmd.Flags().String("modulo", "", "Compute the hash modulo the specified string")
	hash_pathCmd.Flags().Bool("sri", false, "Print the hash in SRI format")
	hash_pathCmd.Flags().String("type", "", "hash algorithm")
	hashCmd.AddCommand(hash_pathCmd)

	common.AddLoggingFlags(hash_pathCmd)

	carapace.Gen(hash_pathCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(hash_pathCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
