package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devenv"
	"github.com/spf13/cobra"
)

var container_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_runCmd).Standalone()

	container_runCmd.Flags().String("copy-args", "", "")

	containerCmd.AddCommand(container_runCmd)

	carapace.Gen(container_runCmd).PositionalCompletion(
		devenv.ActionContainers(),
	)
}
