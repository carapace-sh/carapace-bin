package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_startCmd = &cobra.Command{
	Use:   "start [OPTIONS] CONTAINER [CONTAINER...]",
	Short: "Start one or more stopped containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_startCmd).Standalone()

	container_startCmd.Flags().BoolP("attach", "a", false, "Attach STDOUT/STDERR and forward signals")
	container_startCmd.Flags().String("checkpoint", "", "Restore from this checkpoint")
	container_startCmd.Flags().String("checkpoint-dir", "", "Use a custom checkpoint storage directory")
	container_startCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
	container_startCmd.Flags().BoolP("interactive", "i", false, "Attach container's STDIN")
	containerCmd.AddCommand(container_startCmd)

	carapace.Gen(container_startCmd).FlagCompletion(carapace.ActionMap{
		// TODO checkpoint completion
		"detach-keys": docker.ActionDetachKeys(),
	})
}
