package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var refs_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List references",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_listCmd).Standalone()

	refs_listCmd.Flags().String("color", "", "respect any colors specified in the --format option")
	refs_listCmd.Flags().String("contains", "", "only list refs which contain the specified commit")
	refs_listCmd.Flags().String("count", "", "limit refs")
	refs_listCmd.Flags().StringSlice("exclude", nil, "only list refs which do not match any excluded pattern(s)")
	refs_listCmd.Flags().String("format", "", "a string that interpolates %(fieldname) from a ref being shown")
	refs_listCmd.Flags().Bool("ignore-case", false, "sorting and filtering refs are case insensitive")
	refs_listCmd.Flags().Bool("include-root-refs", false, "also include HEAD ref and pseudorefs")
	refs_listCmd.Flags().String("merged", "", "only list refs whose tips are reachable from the specified commit")
	refs_listCmd.Flags().String("no-contains", "", "only list refs which don't contain the specified commit")
	refs_listCmd.Flags().String("no-merged", "", "only list refs whose tips are not reachable from the specified commit")
	refs_listCmd.Flags().Bool("omit-empty", false, "do not print a newline after formatted refs")
	refs_listCmd.Flags().Bool("perl", false, "quote strings for perl")
	refs_listCmd.Flags().String("points-at", "", "only list refs which points at the given object")
	refs_listCmd.Flags().Bool("python", false, "quote strings for python")
	refs_listCmd.Flags().Bool("shell", false, "quote strings for shell")
	refs_listCmd.Flags().StringSlice("sort", nil, "field name to sort on")
	refs_listCmd.Flags().String("start-after", "", "start iteration after the provided marker")
	refs_listCmd.Flags().Bool("stdin", false, "read reference patterns from stdin")
	refs_listCmd.Flags().Bool("tcl", false, "quote strings for tcl")
	refsCmd.AddCommand(refs_listCmd)

	refs_listCmd.MarkFlagsMutuallyExclusive(
		"shell",
		"perl",
		"python",
		"tcl",
	)

	carapace.Gen(refs_listCmd).FlagCompletion(carapace.ActionMap{
		"color":       carapace.ActionValues("always", "never", "auto").StyleF(style.ForKeyword),
		"contains":    git.ActionRefs(git.RefOption{}.Default()),
		"exclude":     git.ActionRefs(git.RefOption{}.Default()),
		"merged":      git.ActionRefs(git.RefOption{}.Default()),
		"no-contains": git.ActionRefs(git.RefOption{}.Default()),
		"no-merged":   git.ActionRefs(git.RefOption{}.Default()),
		"points-at":   git.ActionRefs(git.RefOption{}.Default()),
		"sort":        git.ActionFieldNames(),
	})

	carapace.Gen(refs_listCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if refs_listCmd.Flag("stdin").Changed {
				return carapace.ActionValues()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)
}
