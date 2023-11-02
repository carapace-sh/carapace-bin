package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var abandonCmd = &cobra.Command{
	Use:   "abandon [OPTIONS] [REVISIONS]...",
	Short: "Abandon a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(abandonCmd).Standalone()

	abandonCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	abandonCmd.Flags().BoolP("summary", "s", false, "Do not print every abandoned commit on a separate line")
	rootCmd.AddCommand(abandonCmd)

	carapace.Gen(abandonCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
