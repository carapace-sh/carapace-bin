package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hash_toBase64Cmd = &cobra.Command{
	Use:   "to-base64",
	Short: "convert a hash to base-64 representation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_toBase64Cmd).Standalone()

	hash_toBase64Cmd.Flags().String("type", "", "hash algorithm")
	hashCmd.AddCommand(hash_toBase64Cmd)

	addLoggingFlags(hash_toBase64Cmd)

	carapace.Gen(hash_fileCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(hash_toBase64Cmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
