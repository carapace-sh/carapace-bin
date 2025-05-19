package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hash_toBase32Cmd = &cobra.Command{
	Use:   "to-base32",
	Short: "convert a hash to base-32 representation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_toBase32Cmd).Standalone()

	hash_toBase32Cmd.Flags().String("type", "", "hash algorithm")
	hashCmd.AddCommand(hash_toBase32Cmd)

	addLoggingFlags(hash_toBase32Cmd)

	carapace.Gen(hash_fileCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(hash_toBase32Cmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
