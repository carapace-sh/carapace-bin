package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var satisfyCmd = &cobra.Command{
	Use:   "satisfy [pattern]...",
	Short: "satisfy dependency strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(satisfyCmd).Standalone()

	rootCmd.AddCommand(satisfyCmd)

	carapace.Gen(satisfyCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch(),
	)
}
