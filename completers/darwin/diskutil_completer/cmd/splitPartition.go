package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var splitPartitionCmd = &cobra.Command{
	Use:   "splitPartition",
	Short: "Split an existing partition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitPartitionCmd).Standalone()
	rootCmd.AddCommand(splitPartitionCmd)
	carapace.Gen(splitPartitionCmd).PositionalCompletion(fs.ActionBlockDevices())
}
