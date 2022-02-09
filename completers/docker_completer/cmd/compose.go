package cmd

import (
	"github.com/rsteube/carapace"
	compose "github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd"
	"github.com/rsteube/carapace-bin/pkg/actions/invoke"
	"github.com/spf13/cobra"
)

var composeCmd = &cobra.Command{
	Use:                "compose",
	Short:              "Define and run multi-container applications with Docker",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(composeCmd).Standalone()

	rootCmd.AddCommand(composeCmd)

	carapace.Gen(composeCmd).PositionalAnyCompletion(
		invoke.ActionInvoke(compose.Execute),
	)
}
