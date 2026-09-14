package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var store_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "perform garbage collection on a Nix store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_gcCmd).Standalone()

	store_gcCmd.Flags().Bool("dry-run", false, "Show what this command would do without doing it")
	store_gcCmd.Flags().String("max", "", "Stop after freeing n bytes of disk space")
	storeCmd.AddCommand(store_gcCmd)

	common.AddLoggingFlags(store_gcCmd)
}
