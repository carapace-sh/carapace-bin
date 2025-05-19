package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hash_toBase16Cmd = &cobra.Command{
	Use:   "to-base16",
	Short: "convert a hash to base-16 representation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_toBase16Cmd).Standalone()

	hash_toBase16Cmd.Flags().String("type", "", "hash algorithm")
	hashCmd.AddCommand(hash_toBase16Cmd)

	addLoggingFlags(hash_toBase16Cmd)

	carapace.Gen(hash_fileCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(hash_toBase16Cmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
