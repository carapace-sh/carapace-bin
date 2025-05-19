package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an OCI image that can run as a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	buildCmd.Flags().String("engine", "docker", "Engine used to build the container: 'docker', 'podman'")
	buildCmd.Flags().String("name", "devbox", "name for the container")
	buildCmd.Flags().Bool("no-cache", false, "Do not use a cache")
	buildCmd.Flags().StringSlice("tags", nil, "tags for the container")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
		"engine": carapace.ActionValues("docker", "podman"),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
