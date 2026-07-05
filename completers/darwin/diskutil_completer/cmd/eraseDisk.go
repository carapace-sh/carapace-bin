package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var eraseDiskCmd = &cobra.Command{
	Use:   "eraseDisk",
	Short: "Erase and format a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eraseDiskCmd).Standalone()
	rootCmd.AddCommand(eraseDiskCmd)
	carapace.Gen(eraseDiskCmd).PositionalCompletion(
		carapace.ActionValues("APFS", "HFS+", "ExFAT", "MS-DOS"),
		fs.ActionBlockDevices(),
	)
}
