package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var forEachRefCmd = &cobra.Command{
	Use:     "for-each-ref",
	Short:   "Output information on each ref",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(forEachRefCmd).Standalone()

	forEachRefCmd.Flags().String("color", "", "respect any colors specified in the --format option")
	forEachRefCmd.Flags().String("contains", "", "only list refs which contain the specified commit")
	forEachRefCmd.Flags().String("count", "", "limit refs")
	forEachRefCmd.Flags().StringSlice("exclude", nil, "only list refs which do not match any excluded pattern(s)")
	forEachRefCmd.Flags().String("format", "", "a string that interpolates %(fieldname) from a ref being shown")
	forEachRefCmd.Flags().Bool("ignore-case", false, "sorting and filtering refs are case insensitive")
	forEachRefCmd.Flags().String("merged", "", "only list refs whose tips are reachable from the specified commit")
	forEachRefCmd.Flags().String("no-contains", "", "only list refs which don’t contain the specified commit")
	forEachRefCmd.Flags().String("no-merged", "", "only list refs whose tips are not reachable from the specified commit")
	forEachRefCmd.Flags().Bool("omit-empty", false, "do not print a newline after formatted refs")
	forEachRefCmd.Flags().Bool("perl", false, "quote strings for perl")
	forEachRefCmd.Flags().String("points-at", "", "only list refs which points at the given object")
	forEachRefCmd.Flags().Bool("python", false, "quote strings for python")
	forEachRefCmd.Flags().Bool("shell", false, "quote strings for shell ")
	forEachRefCmd.Flags().StringSlice("sort", nil, "field name to sort on")
	forEachRefCmd.Flags().Bool("stdin", false, "read patterns from standard input")
	forEachRefCmd.Flags().Bool("tcl", false, "quote strings for tcl")
	rootCmd.AddCommand(forEachRefCmd)

	forEachRefCmd.MarkFlagsMutuallyExclusive(
		"shell",
		"perl",
		"python",
		"tcl",
	)

	carapace.Gen(forEachRefCmd).FlagCompletion(carapace.ActionMap{
		"color":       carapace.ActionValues("always", "never", "auto").StyleF(style.ForKeyword),
		"contains":    git.ActionRefs(git.RefOption{}.Default()),
		"exclude":     git.ActionRefs(git.RefOption{}.Default()),
		"merged":      git.ActionRefs(git.RefOption{}.Default()),
		"no-contains": git.ActionRefs(git.RefOption{}.Default()),
		"no-merged":   git.ActionRefs(git.RefOption{}.Default()),
		"points-at":   git.ActionRefs(git.RefOption{}.Default()),
		"sort":        git.ActionFieldNames(),
	})

	carapace.Gen(forEachRefCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if forEachRefCmd.Flag("stdin").Changed {
				return carapace.ActionValues()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)
}
