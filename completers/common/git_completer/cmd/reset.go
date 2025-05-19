package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Reset current HEAD to the specified state",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().Bool("hard", false, "reset HEAD, index and working tree")
	resetCmd.Flags().BoolP("intent-to-add", "N", false, "record only the fact that removed paths will be added later")
	resetCmd.Flags().Bool("keep", false, "reset HEAD but keep local changes")
	resetCmd.Flags().Bool("merge", false, "reset HEAD, index and working tree")
	resetCmd.Flags().Bool("mixed", false, "reset HEAD and index")
	resetCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	resetCmd.Flags().Bool("pathspec-file-nul", false, "pathspec elements are separated with NUL character")
	resetCmd.Flags().String("pathspec-from-file", "", "read pathspec from file")
	resetCmd.Flags().BoolP("quiet", "q", false, "be quiet, only report errors")
	resetCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	resetCmd.Flags().Bool("soft", false, "reset only HEAD")
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flag("recurse-submodules").NoOptDefVal = " "

	modeFlags := []string{"soft", "mixed", "hard", "merge", "keep"}
	resetCmd.MarkFlagsMutuallyExclusive(
		append(modeFlags, "patch")..., // TODO verify - either one of the "modes" above or patch allowed
	)

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"pathspec-from-file": carapace.ActionFiles(),
	})

	carapace.Gen(resetCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, ".") {
				for _, name := range modeFlags {
					if f := resetCmd.Flag(name); f != nil && f.Changed {
						return carapace.ActionValues() // complete only commits for these flags
					}
				}
				return git.ActionRefDiffs()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)

	carapace.Gen(resetCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			toFilter := make([]string, 0)
			for _, arg := range c.Args {
				// TODO directories? globs possible? should work for most cases though
				switch {
				case !strings.HasPrefix(c.Value, "."):
					toFilter = append(toFilter, strings.TrimPrefix(arg, "./"))
				case !strings.HasPrefix(arg, "."):
					toFilter = append(toFilter, "./"+arg)
				default:
					toFilter = append(toFilter, arg)
				}
			}

			// multipart completion here as there can be a lot of differences
			switch {
			case strings.HasPrefix(c.Args[0], "."):
				return git.ActionRefDiffs().Filter(toFilter...).MultiParts("/").StyleF(style.ForPathExt)
			default:
				return git.ActionRefDiffs(c.Args[0]).Filter(toFilter[1:]...).MultiParts("/").StyleF(style.ForPathExt)
			}
		}),
	)

	carapace.Gen(resetCmd).DashAnyCompletion(
		carapace.ActionPositional(resetCmd), // TODO should only be files, but will suffice for now
	)
}
