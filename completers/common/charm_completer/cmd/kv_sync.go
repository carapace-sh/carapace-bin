package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var kv_syncCmd = &cobra.Command{
	Use:   "sync [@DB]",
	Short: "Sync local db with latest Charm Cloud db.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_syncCmd).Standalone()

	kvCmd.AddCommand(kv_syncCmd)

	carapace.Gen(kv_syncCmd).PositionalCompletion(
		charm.ActionDatabases(),
	)
}
