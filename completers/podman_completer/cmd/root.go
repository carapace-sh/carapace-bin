package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/cobracomplete"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "podman",
	Short:              "Simple management tool for pods, containers and images",
	Long:               "https://podman.io/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		// TODO just bridging it for now
		cobracomplete.ActionCobraComplete("podman"),
	)
}
