package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build an OCI image that can run as a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	buildCmd.Flags().String("engine", "docker", "Engine used to build the container: 'docker', 'podman'")
	buildCmd.Flags().String("name", "devbox", "name for the container")
	buildCmd.Flags().Bool("no-cache", false, "Do not use a cache")
	buildCmd.Flags().StringSlice("tags", []string{}, "tags for the container")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"engine": carapace.ActionValues("docker", "podman"),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
