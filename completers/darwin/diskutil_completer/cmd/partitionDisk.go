package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var partitionDiskCmd = &cobra.Command{
	Use:   "partitionDisk",
	Short: "Partition a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partitionDiskCmd).Standalone()
	rootCmd.AddCommand(partitionDiskCmd)
	carapace.Gen(partitionDiskCmd).PositionalCompletion(
		carapace.ActionValues("APFS", "HFS+", "ExFAT", "MS-DOS"),
		fs.ActionBlockDevices(),
	)
}
