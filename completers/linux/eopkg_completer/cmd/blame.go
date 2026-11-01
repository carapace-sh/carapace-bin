package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:     "blame <packagename>",
	Aliases: []string{"bl"},
	Short:   "show history entry for a given package to show the package's changelog",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	blameCmd.Flags().BoolP("all", "a", false, "show blame for the entire history of the package")
	blameCmd.Flags().IntP("release", "r", 0, "only show blame for the given release number")

	rootCmd.AddCommand(blameCmd)

	carapace.Gen(blameCmd).PositionalAnyCompletion(
		eopkg.ActionPackages().FilterArgs(),
	)
}
