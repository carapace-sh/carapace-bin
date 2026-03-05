package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var arrangeCmd = &cobra.Command{
	Use:   "arrange",
	Short: "Interactively arrange the commit graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arrangeCmd).Standalone()

	arrangeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	arrangeCmd.Flags().StringSliceP("revisions", "r", []string{"@ | ancestors(immutable_heads().., 2) | heads(immutable_heads())"}, "The revisions to edit")
	rootCmd.AddCommand(arrangeCmd)

	carapace.Gen(arrangeCmd).FlagCompletion(carapace.ActionMap{
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
