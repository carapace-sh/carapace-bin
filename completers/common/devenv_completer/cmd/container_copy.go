package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devenv"
	"github.com/spf13/cobra"
)

var container_copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a container to registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_copyCmd).Standalone()

	container_copyCmd.Flags().String("copy-args", "", "")
	container_copyCmd.Flags().StringP("registry", "r", "", "")

	containerCmd.AddCommand(container_copyCmd)

	carapace.Gen(container_copyCmd).PositionalCompletion(
		devenv.ActionContainers(),
	)
}
