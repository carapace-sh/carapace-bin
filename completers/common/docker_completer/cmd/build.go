package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:                "build [OPTIONS] PATH | URL | -",
	Short:              "Build an image from a Dockerfile",
	Run:                func(cmd *cobra.Command, args []string) {},
	GroupID:            "common",
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("docker", "buildx", "build"),
	)
}
