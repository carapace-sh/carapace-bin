package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var madisonCmd = &cobra.Command{
	Use:   "madison",
	Short: "mimic the output format of the Debian madison tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(madisonCmd).Standalone()

	rootCmd.AddCommand(madisonCmd)

	carapace.Gen(madisonCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
