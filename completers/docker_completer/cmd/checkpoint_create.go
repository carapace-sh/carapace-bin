package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var checkpoint_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] CONTAINER CHECKPOINT",
	Short: "Create a checkpoint from a running container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkpoint_createCmd).Standalone()

	checkpoint_createCmd.Flags().String("checkpoint-dir", "", "Use a custom checkpoint storage directory")
	checkpoint_createCmd.Flags().Bool("leave-running", false, "Leave the container running after checkpoint")
	checkpointCmd.AddCommand(checkpoint_createCmd)

	carapace.Gen(checkpoint_createCmd).FlagCompletion(carapace.ActionMap{
		"checkpoint-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(checkpoint_createCmd).PositionalCompletion(
		docker.ActionContainers(),
	)
}
