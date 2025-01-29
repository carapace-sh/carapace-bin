package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_rmCmd = &cobra.Command{
	Use:     "rm [options] VOLUME [VOLUME...]",
	Short:   "Remove one or more volumes",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_rmCmd).Standalone()

	volume_rmCmd.Flags().BoolP("all", "a", false, "Remove all volumes")
	volume_rmCmd.Flags().BoolP("force", "f", false, "Remove a volume by force, even if it is being used by a container")
	volume_rmCmd.Flags().StringP("time", "t", "", "Seconds to wait for running containers to stop before killing the container")
	volumeCmd.AddCommand(volume_rmCmd)
}
