package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge is identical to remove except that packages are removed and purged",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	rootCmd.AddCommand(purgeCmd)

	carapace.Gen(purgeCmd).PositionalAnyCompletion(
		apt.ActionPackages().FilterArgs(),
	)
}
