package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var resilverCmd = &cobra.Command{
	Use:     "resilver pool...",
	Short:   "start a resilver",
	GroupID: "fault",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resilverCmd).Standalone()

	rootCmd.AddCommand(resilverCmd)

	carapace.Gen(resilverCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
