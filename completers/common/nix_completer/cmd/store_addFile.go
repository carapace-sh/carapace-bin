package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var store_addFileCmd = &cobra.Command{
	Use:   "add-file",
	Short: "add a regular file to the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_addFileCmd).Standalone()

	store_addFileCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	store_addFileCmd.Flags().String("hash-algo", "", "Hash algorithm (blake3, md5, sha1, sha256, or sha512)")
	store_addFileCmd.Flags().String("mode", "", "How to compute the content-address of the store object")
	store_addFileCmd.Flags().StringP("name", "n", "", "Override the name component of the store path. It defaults to the base name of path")
	storeCmd.AddCommand(store_addFileCmd)

	common.AddLoggingFlags(store_addFileCmd)

	carapace.Gen(store_addFileCmd).FlagCompletion(carapace.ActionMap{
		"hash-algo": carapace.ActionValues("blake3", "md5", "sha1", "sha256", "sha512"),
		"mode":      carapace.ActionValues("nar", "flat", "git"),
	})

	carapace.Gen(store_addFileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
