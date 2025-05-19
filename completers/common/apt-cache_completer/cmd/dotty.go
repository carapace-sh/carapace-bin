package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var dottyCmd = &cobra.Command{
	Use:   "dotty",
	Short: "generates output suitable for use by dotty",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dottyCmd).Standalone()

	rootCmd.AddCommand(dottyCmd)

	carapace.Gen(dottyCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
