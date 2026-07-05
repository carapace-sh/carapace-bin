package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var verifyVolumeCmd = &cobra.Command{
	Use:   "verifyVolume",
	Short: "Verify the filesystem of a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyVolumeCmd).Standalone()
	rootCmd.AddCommand(verifyVolumeCmd)
	carapace.Gen(verifyVolumeCmd).PositionalCompletion(fs.ActionBlockDevices())
}
