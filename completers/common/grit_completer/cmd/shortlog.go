package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var shortlogCmd = &cobra.Command{
	Use:     "shortlog",
	Short:   "List the commits on this branch that aren't on the target branch yet",
	Aliases: []string{"sl", "ahead"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shortlogCmd).Standalone()

	shortlogCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	shortlogCmd.Flags().String("target", "", "Compare against this ref or revision instead of the discovered target branch")
	rootCmd.AddCommand(shortlogCmd)

	carapace.Gen(shortlogCmd).FlagCompletion(carapace.ActionMap{
		"target": git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}),
	})
}
