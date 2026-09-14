package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var hash_convertCmd = &cobra.Command{
	Use:   "convert [flags] [hashes...]",
	Short: "convert between hash formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hash_convertCmd).Standalone()

	hash_convertCmd.Flags().String("from", "", "Hash format (base16, nix32, base64, sri)")
	hash_convertCmd.Flags().String("hash-algo", "", "Hash algorithm (md5, sha1, sha256, sha512)")
	hash_convertCmd.Flags().String("to", "", "Hash format (base16, nix32, base64, sri). Default: sri")
	hashCmd.AddCommand(hash_convertCmd)

	common.AddLoggingFlags(hash_convertCmd)

	carapace.Gen(hash_convertCmd).FlagCompletion(carapace.ActionMap{
		"from":      carapace.ActionValues("base16", "nix32", "base64", "sri"),
		"hash-algo": carapace.ActionValues("blake3", "md5", "sha1", "sha256", "sha512"),
		"to":        carapace.ActionValues("base16", "nix32", "base64", "sri"),
	})
}
