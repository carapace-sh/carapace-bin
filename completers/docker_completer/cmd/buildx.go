package cmd

import (
	"github.com/rsteube/carapace"
	buildx "github.com/rsteube/carapace-bin/completers/docker-buildx_completer/cmd"
	"github.com/spf13/cobra"
)

var buildxCmd = &cobra.Command{
	Use:                "buildx",
	Short:              "Extended build capabilities with BuildKit",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(buildxCmd).Standalone()

	rootCmd.AddCommand(buildxCmd)

	carapace.Gen(buildxCmd).PositionalAnyCompletion(
		carapace.ActionInvoke(buildx.Execute),
	)
}
