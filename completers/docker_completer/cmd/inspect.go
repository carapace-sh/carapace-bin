package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:    "inspect [OPTIONS] NAME|ID [NAME|ID...]",
	Short:  "Return low-level information on Docker objects",
	Hidden: os.Getenv("DOCKER_HIDE_LEGACY_COMMANDS") == "1",
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()
	inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	inspectCmd.Flags().BoolP("size", "s", false, "Display total file sizes if the type is container")
	inspectCmd.Flags().String("type", "", "Return JSON for specified type")
	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).PositionalAnyCompletion(
		carapace.Batch(
			docker.ActionContainers(),
			docker.ActionNetworks(),
			docker.ActionNodes().Suppress("This node is not a swarm manager"),
			docker.ActionRepositoryTags(),
			docker.ActionSecrets().Suppress("This node is not a swarm manager"),
			docker.ActionServices().Suppress("This node is not a swarm manager"),
			docker.ActionVolumes(),
		).ToA(),
	)
}
