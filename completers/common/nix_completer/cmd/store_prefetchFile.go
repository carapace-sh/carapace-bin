package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var store_prefetchFileCmd = &cobra.Command{
	Use:   "prefetch-file",
	Short: "download a file into the Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_prefetchFileCmd).Standalone()

	store_prefetchFileCmd.Flags().Bool("executable", false, "Make the resulting file executable")
	store_prefetchFileCmd.Flags().String("expected-hash", "", "The expected hash of the file")
	store_prefetchFileCmd.Flags().String("hash-type", "", "hash algorithm")
	store_prefetchFileCmd.Flags().String("name", "", "Override the name component of the resulting store path")
	store_prefetchFileCmd.Flags().Bool("unpack", false, "Unpack the archive (which must be a tarball or zip file) and add the unpacked contents to the store")
	storeCmd.AddCommand(store_prefetchFileCmd)

	common.AddLoggingFlags(store_prefetchFileCmd)

	carapace.Gen(store_prefetchFileCmd).FlagCompletion(carapace.ActionMap{
		"hash-type": carapace.ActionValues("md5", "sha1", "sha256", "sha512"),
	})
}
