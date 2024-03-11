package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "source causes apt-get to fetch source packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sourceCmd).Standalone()

	rootCmd.AddCommand(sourceCmd)

	carapace.Gen(sourceCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch().FilterArgs(),
	)
}
