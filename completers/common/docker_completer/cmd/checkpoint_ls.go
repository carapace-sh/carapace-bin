package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var checkpoint_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS] CONTAINER",
	Short:   "List checkpoints for a container",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkpoint_lsCmd).Standalone()

	checkpoint_lsCmd.Flags().String("checkpoint-dir", "", "Use a custom checkpoint storage directory")
	checkpointCmd.AddCommand(checkpoint_lsCmd)

	carapace.Gen(checkpoint_lsCmd).FlagCompletion(carapace.ActionMap{
		"checkpoint-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(checkpoint_lsCmd).PositionalCompletion(
		docker.ActionContainers(),
	)
}
