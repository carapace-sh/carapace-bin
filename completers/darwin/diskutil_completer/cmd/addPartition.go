package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var addPartitionCmd = &cobra.Command{
	Use:   "addPartition",
	Short: "Add a new partition to a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addPartitionCmd).Standalone()
	rootCmd.AddCommand(addPartitionCmd)
	carapace.Gen(addPartitionCmd).PositionalCompletion(fs.ActionBlockDevices())
}
