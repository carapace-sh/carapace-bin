package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var alpha_generateCmd = &cobra.Command{
	Use:   "generate [OPTIONS] [CONTAINERS...]",
	Short: "EXPERIMENTAL - Generate a Compose file from existing containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_generateCmd).Standalone()

	alpha_generateCmd.Flags().String("format", "", "Format the output. Values: [yaml | json]")
	alpha_generateCmd.Flags().String("name", "", "Project name to set in the Compose file")
	alpha_generateCmd.Flags().String("project-dir", "", "Directory to use for the project")
	alphaCmd.AddCommand(alpha_generateCmd)

	carapace.Gen(alpha_generateCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("yaml", "json"),
		"project-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(alpha_generateCmd).PositionalAnyCompletion(
		docker.ActionContainers().FilterArgs(),
	)
}
