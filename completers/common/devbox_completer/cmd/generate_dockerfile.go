package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Generate a Dockerfile that replicates devbox shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_dockerfileCmd).Standalone()

	generate_dockerfileCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	generate_dockerfileCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	generate_dockerfileCmd.Flags().String("for", "", "Generate Dockerfile for a specific type of container (dev, prod)")
	generate_dockerfileCmd.Flags().BoolP("force", "f", false, "force overwrite existing files")
	generate_dockerfileCmd.Flags().Bool("root-user", false, "Use root as default user inside the container")
	generate_dockerfileCmd.Flag("for").Hidden = true
	generateCmd.AddCommand(generate_dockerfileCmd)

	// TODO environment
	carapace.Gen(generate_dockerfileCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
		"for":    carapace.ActionValues("dev", "prod"),
	})
}
