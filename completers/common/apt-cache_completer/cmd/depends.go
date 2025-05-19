package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:   "depends",
	Short: "shows a listing of each dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()

	rootCmd.AddCommand(dependsCmd)

	carapace.Gen(dependsCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
