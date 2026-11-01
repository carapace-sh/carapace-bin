package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var hash_toSriCmd = &cobra.Command{
	Use:   "to-sri",
	Short: "convert a hash to SRI representation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_toSriCmd).Standalone()

	hash_toSriCmd.Flags().String("type", "", "hash algorithm")
	hashCmd.AddCommand(hash_toSriCmd)

	common.AddLoggingFlags(hash_toSriCmd)

	carapace.Gen(hash_fileCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})

	carapace.Gen(hash_toSriCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
