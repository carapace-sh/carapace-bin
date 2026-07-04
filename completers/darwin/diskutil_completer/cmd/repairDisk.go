package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var repairDiskCmd = &cobra.Command{
	Use:   "repairDisk",
	Short: "Repair the partition map of a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repairDiskCmd).Standalone()
	rootCmd.AddCommand(repairDiskCmd)
	carapace.Gen(repairDiskCmd).PositionalCompletion(fs.ActionBlockDevices())
}
