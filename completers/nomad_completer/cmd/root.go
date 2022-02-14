package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/posenercomplete"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "nomad",
	Short:              "https://github.com/hashicorp/nomad",
	Long:               "https://www.nomadproject.io/",
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
		posenercomplete.ActionPosenerComplete("nomad"),
	)
}
