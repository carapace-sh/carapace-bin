package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
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
	abandonCmd.Flags().Bool("restore-descendants", false, "Do not modify the content of the children of the abandoned commits")
	abandonCmd.Flags().Bool("retain-bookmarks", false, "Do not delete bookmarks pointing to the revisions to abandon")
	rootCmd.AddCommand(abandonCmd)

	carapace.Gen(abandonCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
