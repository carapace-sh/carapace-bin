package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] VOLUME [VOLUME...]",
	Short:   "Remove one or more volumes",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_rmCmd).Standalone()

	volume_rmCmd.Flags().BoolP("force", "f", false, "Force the removal of one or more volumes")
	volumeCmd.AddCommand(volume_rmCmd)

	carapace.Gen(volume_rmCmd).PositionalAnyCompletion(docker.ActionVolumes())
}
