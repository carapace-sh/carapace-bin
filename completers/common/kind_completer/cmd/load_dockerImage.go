package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var load_dockerImageCmd = &cobra.Command{
	Use:   "docker-image <IMAGE> [IMAGE...]",
	Short: "Loads docker images from host into nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(load_dockerImageCmd).Standalone()

	load_dockerImageCmd.Flags().StringP("name", "n", "kind", "the cluster context name")
	load_dockerImageCmd.Flags().StringSlice("nodes", nil, "comma separated list of nodes to load images into")

	carapace.Gen(load_dockerImageCmd).FlagCompletion(carapace.ActionMap{
		"name":  action.ActionClusterNames(),
		"nodes": action.ActionNodes(load_dockerImageCmd),
	})

	carapace.Gen(load_dockerImageCmd).PositionalAnyCompletion(
		docker.ActionRepositoryTags(),
	)

	loadCmd.AddCommand(load_dockerImageCmd)
}
