package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsListSnapshotsCmd = &cobra.Command{
	Use:   "listSnapshots",
	Short: "Show APFS snapshots on a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsListSnapshotsCmd).Standalone()
	apfsCmd.AddCommand(apfsListSnapshotsCmd)
	carapace.Gen(apfsListSnapshotsCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
