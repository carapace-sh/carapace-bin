package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var reflogCmd = &cobra.Command{
	Use:     "reflog",
	Short:   "Show, delete, or expire reflog entries",
	GroupID: "history",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflogCmd).Standalone()

	reflogCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	reflogCmd.Flags().String("limit", "", "Show at most this many entries (newest first)")
	rootCmd.AddCommand(reflogCmd)

	carapace.Gen(reflogCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
