package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:                "build",
	Short:              "Builds and analyzes a docker image from a Dockerfile (this is a thin wrapper for the `docker build` command).",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(buildCmd).Standalone()
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = append([]string{"build"}, c.Args...)
			return bridge.ActionCarapaceBin("docker").Invoke(c).ToA()
		}),
	)
}
