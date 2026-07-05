package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var verifyDiskCmd = &cobra.Command{
	Use:   "verifyDisk",
	Short: "Verify the partition map of a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyDiskCmd).Standalone()
	rootCmd.AddCommand(verifyDiskCmd)
	carapace.Gen(verifyDiskCmd).PositionalCompletion(fs.ActionBlockDevices())
}
