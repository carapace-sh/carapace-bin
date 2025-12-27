package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "displays the package records for the named packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
