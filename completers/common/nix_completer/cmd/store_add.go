package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var store_addCmd = &cobra.Command{
	Use:   "add [path]",
	Short: "add a file or directory to the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_addCmd).Standalone()

	store_addCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	store_addCmd.Flags().String("hash-algo", "", "Hash algorithm (blake3, md5, sha1, sha256, or sha512)")
	store_addCmd.Flags().String("mode", "", "How to compute the content-address of the store object")
	store_addCmd.Flags().StringP("name", "n", "", "Override the name component of the store path. It defaults to the base name of path")
	storeCmd.AddCommand(store_addCmd)

	common.AddLoggingFlags(store_addCmd)

	carapace.Gen(store_addCmd).FlagCompletion(carapace.ActionMap{
		"hash-algo": carapace.ActionValues("blake3", "md5", "sha1", "sha256", "sha512"),
		"mode":      carapace.ActionValues("nar", "flat", "git"),
	})

	carapace.Gen(store_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
