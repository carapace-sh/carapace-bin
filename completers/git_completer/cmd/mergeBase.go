package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeBaseCmd = &cobra.Command{
	Use:     "merge-base",
	Short:   "Find as good common ancestors as possible for a merge",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(mergeBaseCmd).Standalone()

	mergeBaseCmd.Flags().BoolP("all", "a", false, "output all merge bases for the commits")
	mergeBaseCmd.Flags().Bool("fork-point", false, "find the point at which a branch forked from another branch <ref>")
	mergeBaseCmd.Flags().Bool("independent", false, "print a minimal subset of the supplied commits with the same ancestors")
	mergeBaseCmd.Flags().Bool("is-ancestor", false, "check if the first <commit> is an ancestor of the second <commit>")
	mergeBaseCmd.Flags().Bool("octopus", false, "compute the best common ancestors of all supplied commits")
	rootCmd.AddCommand(mergeBaseCmd)

	mergeBaseCmd.MarkFlagsMutuallyExclusive(
		"fork-point",
		"independent",
		"is-ancestor",
		"octopus",
	)

	carapace.Gen(mergeBaseCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch {
			case mergeBaseCmd.Flag("is-ancestor").Changed && len(c.Args) > 1,
				mergeBaseCmd.Flag("fork-point").Changed && len(c.Args) > 1:
				return carapace.ActionValues()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)
}
