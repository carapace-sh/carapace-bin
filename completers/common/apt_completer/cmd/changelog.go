package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var changelogCmd = &cobra.Command{
	Use:   "changelog [pattern]...",
	Short: "show package changelog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(changelogCmd).Standalone()
	rootCmd.AddCommand(changelogCmd)

	carapace.Gen(changelogCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch(),
	)
}
