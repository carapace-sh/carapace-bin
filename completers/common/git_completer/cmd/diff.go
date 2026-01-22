package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:     "diff",
	Short:   "Show changes between commits",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	// TODO move into common flag groups
	diffCmd.Flags().StringS("G", "G", "", "Look for differences whose patch text contains added/removed lines that match <regex>")
	diffCmd.Flags().String("follow", "", "Continue listing the history of a file beyond renames")
	diffCmd.Flags().StringArrayP("ignore-matching-lines", "I", nil, "Ignore changes whose all lines match <regex>")
	diffCmd.Flags().Bool("quiet", false, "Disable all output of the program")
	diffCmd.Flags().String("rotate-to", "", "Move the files before the named <file> to the end")
	diffCmd.Flags().String("skip-to", "", "Discard the files before the named <file> from the output")
	diffCmd.Flags().Bool("staged", false, "Show diff between index and named commit")
	common.AddDiffFlags(diffCmd)
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"follow":  carapace.ActionFiles(), // TODO complete files of specific revision/modified between commits?
		"skip-to": carapace.ActionFiles(), // TODO complete files of specific revision/modified between commits?
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		actionDiffArgs(diffCmd),
	)

	carapace.Gen(diffCmd).DashAnyCompletion(
		carapace.ActionPositional(diffCmd),
	)
}

func actionDiffArgs(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if cmd.Flag("no-index").Changed {
			switch len(c.Args) {
			case 0, 1:
				return carapace.ActionFiles().FilterArgs()
			default:
				return carapace.ActionValues()
			}
		}

		batch := carapace.Batch()

		filtered := make([]string, 0)
		for index, arg := range c.Args {
			if index == 0 && strings.Contains(arg, "..") { // assume refrange - TODO what about '{ref}...{ref}'
				filtered = append(filtered, arg)
				break
			}

			if err := c.Command("git", "rev-parse", "--verify", arg).Run(); err != nil {
				break
			}
			filtered = append(filtered, arg)
		}

		if len(filtered) == len(c.Args) && cmd.Flags().ArgsLenAtDash() < 0 {
			switch len(filtered) {
			case 0:
				batch = append(batch, git.ActionRefRanges(git.RefOption{}.Default()))
			default:
				switch {
				case strings.Contains(c.Args[0], ".."): // skip if we already have a refrange
				case cmd.Flag("cached").Changed: // skip as '-cached' accepts only on ref
				default:
					batch = append(batch, git.ActionRefs(git.RefOption{}.Default()))
				}
			}
		}

		expanded := filtered
		if len(filtered) > 0 {
			// TODO support/suppress '{ref}...{ref}'??
			expanded = append(strings.SplitN(filtered[0], "..", 2), filtered[1:]...) // split refrange if any
		}

		switch len(expanded) {
		case 0:
			if !cmd.Flag("cached").Changed {
				// TODO `git diff` fails on deleted files (seems still worth seeing them in completion though)
				batch = append(batch, git.ActionChanges(git.ChangeOpts{Unstaged: true}).Filter(cmd.Flags().Args()...))
			}
		default:
			// TODO handle --merge-base with more than 2 refs
			var action carapace.Action
			if cmd.Flag("cached").Changed {
				if len(expanded) > 0 {
					action = git.ActionCachedDiffs(expanded[0])
				}
			} else {
				action = git.ActionRefDiffs(expanded...)
			}

			if len(expanded) > 0 { // multipart for potentially large diffs
				action = action.MultiParts("/").StyleF(style.ForPathExt).Tag("changed files")
			}
			batch = append(batch, action.Filter(cmd.Flags().Args()[len(filtered):]...))
		}

		return batch.ToA()
	})
}
