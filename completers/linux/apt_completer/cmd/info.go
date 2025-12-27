package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info package...",
	Short: "show package information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch().FilterArgs(),
	)
}
