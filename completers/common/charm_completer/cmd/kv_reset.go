package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var kv_resetCmd = &cobra.Command{
	Use:   "reset [@DB]",
	Short: "Delete local db and pull down fresh copy from Charm Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_resetCmd).Standalone()

	kvCmd.AddCommand(kv_resetCmd)

	carapace.Gen(kv_resetCmd).PositionalCompletion(
		charm.ActionDatabases(),
	)
}
