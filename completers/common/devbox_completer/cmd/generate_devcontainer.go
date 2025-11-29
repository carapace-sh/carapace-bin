package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_devcontainerCmd = &cobra.Command{
	Use:   "devcontainer",
	Short: "Generate Dockerfile and devcontainer.json files under .devcontainer/ directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_devcontainerCmd).Standalone()

	generate_devcontainerCmd.Flags().BoolP("force", "f", false, "force overwrite on existing files")
	generate_devcontainerCmd.Flags().Bool("root-user", false, "Use root as default user inside the container")
	generateCmd.AddCommand(generate_devcontainerCmd)
}
