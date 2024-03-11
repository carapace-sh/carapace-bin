package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var checkpoint_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] CONTAINER CHECKPOINT",
	Short:   "Remove a checkpoint",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkpoint_rmCmd).Standalone()

	checkpoint_rmCmd.Flags().String("checkpoint-dir", "", "Use a custom checkpoint storage directory")
	checkpointCmd.AddCommand(checkpoint_rmCmd)

	carapace.Gen(checkpoint_rmCmd).FlagCompletion(carapace.ActionMap{
		"checkpoint-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(checkpoint_rmCmd).PositionalCompletion(
		docker.ActionContainers(),
		// TODO checkpoint completion
	)
}
