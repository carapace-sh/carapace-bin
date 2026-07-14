package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:     "branch",
	Short:   "List, create, or delete branches",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().BoolS("C", "C", false, "shortcut for --copy --force")
	branchCmd.Flags().BoolS("D", "D", false, "shortcut for --delete --force")
	branchCmd.Flags().BoolS("M", "M", false, "shortcut for --move --force")
	branchCmd.Flags().String("abbrev", "", "alter the sha1s minimum display length in the output listing")
	branchCmd.Flags().BoolP("all", "a", false, "list both remote-tracking branches and local branches")
	branchCmd.Flags().String("color", "", "color branches to highlight current, local, and remote-tracking branches")
	branchCmd.Flags().Bool("column", false, "display branch listing in columns")
	branchCmd.Flags().String("contains", "", "only list branches which contain the specified commit (HEAD if not specified)")
	branchCmd.Flags().BoolP("copy", "c", false, "copy a branch together with its config and reflog")
	branchCmd.Flags().Bool("create-reflog", false, "create the branch's reflog")
	branchCmd.Flags().BoolP("delete", "d", false, "delete a branch")
	branchCmd.Flags().Bool("edit-description", false, "open an editor and edit the text to explain what the branch is for")
	branchCmd.Flags().BoolP("force", "f", false, "reset <branchname> to <startpoint>, even if <branchname> exists already")
	branchCmd.Flags().String("format", "", "a string that interpolates %(fieldname) from a branch ref being shown and the object it points at")
	branchCmd.Flags().BoolP("ignore-case", "i", false, "sorting and filtering branches are case insensitive")
	branchCmd.Flags().BoolP("list", "l", false, "list branches")
	branchCmd.Flags().String("merged", "", "only list branches whose tips are reachable from the specified commit (HEAD if not specified)")
	branchCmd.Flags().BoolP("move", "m", false, "move/rename a branch together with its config and reflog")
	branchCmd.Flags().Bool("no-abbrev", false, "display the full sha1s in the output listing rather than abbreviating them")
	branchCmd.Flags().Bool("no-color", false, "turn off branch colors, even when the configuration file gives the default to color output")
	branchCmd.Flags().Bool("no-column", false, "do not display branch listing in columns")
	branchCmd.Flags().String("no-contains", "", "only list branches which don't contain the specified commit (HEAD if not specified)")
	branchCmd.Flags().String("no-merged", "", "only list branches whose tips are not reachable from the specified commit (HEAD if not specified)")
	branchCmd.Flags().Bool("no-track", false, "do not set up upstream configuration, even if branch.autoSetupMerge is true")
	branchCmd.Flags().Bool("omit-empty", false, "do not print a newline after formatted refs")
	branchCmd.Flags().String("points-at", "", "only list branches of the given object")
	branchCmd.Flags().BoolP("quiet", "q", false, "be more quiet when creating or deleting a branch, suppressing non-error messages")
	branchCmd.Flags().Bool("recurse-submodules", false, "cause the current command to recurse into submodules")
	branchCmd.Flags().BoolP("remotes", "r", false, "list or delete (if used with -d) the remote-tracking branches")
	branchCmd.Flags().String("set-upstream-to", "", "set up <branchname>'s tracking information so <upstream> is considered <branchname>'s upstream branch")
	branchCmd.Flags().Bool("show-current", false, "print the name of the current branch")
	branchCmd.Flags().StringArray("sort", nil, "sort based on the key given")
	branchCmd.Flags().StringP("track", "t", "", "when creating a new branch, set up branch.<name>.remote and branch.<name>.merge configuration entries to mark the start-point branch as upstream")
	branchCmd.Flags().Bool("unset-upstream", false, "remove the upstream information for <branchname>")
	branchCmd.Flags().CountP("verbose", "v", "verbose output")
	rootCmd.AddCommand(branchCmd)

	branchCmd.Flag("track").NoOptDefVal = " "

	carapace.Gen(branchCmd).FlagCompletion(carapace.ActionMap{
		"D":               git.ActionRefs(git.RefOption{}.Default()),
		"color":           git.ActionColorModes(),
		"contains":        git.ActionRefs(git.RefOption{}.Default()),
		"delete":          git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}),
		"merged":          git.ActionRefs(git.RefOption{}.Default()),
		"no-contains":     git.ActionRefs(git.RefOption{}.Default()),
		"no-merged":       git.ActionRefs(git.RefOption{}.Default()),
		"points-at":       git.ActionRefs(git.RefOption{RemoteBranches: true, Tags: true}),
		"set-upstream-to": git.ActionRefs(git.RefOption{RemoteBranches: true, Tags: true}),
		"sort":            git.ActionFieldNames(),
		"track":           carapace.ActionValues("direct", "inherit"),
	})

	carapace.Gen(branchCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Args) {
			case 0:
				if branchCmd.Flag("set-upstream-to").Changed ||
					branchCmd.Flag("track").Changed ||
					branchCmd.Flag("no-track").Changed ||
					branchCmd.Flag("unset-upstream").Changed ||
					branchCmd.Flag("M").Changed ||
					branchCmd.Flag("C").Changed ||
					branchCmd.Flag("D").Changed ||
					branchCmd.Flag("move").Changed ||
					branchCmd.Flag("copy").Changed ||
					branchCmd.Flag("delete").Changed ||
					branchCmd.Flag("edit-description").Changed {
					return git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}).FilterArgs()
				}
			case 1:
				if branchCmd.Flag("M").Changed ||
					branchCmd.Flag("C").Changed ||
					branchCmd.Flag("D").Changed ||
					branchCmd.Flag("move").Changed ||
					branchCmd.Flag("copy").Changed ||
					branchCmd.Flag("delete").Changed {
					return git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}).FilterArgs()
				}
			default:
				if branchCmd.Flag("D").Changed ||
					branchCmd.Flag("delete").Changed {
					return git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}).FilterArgs()
				}
			}
			return carapace.ActionValues()
		}),
	)

	carapace.Gen(branchCmd).DashAnyCompletion(
		carapace.ActionPositional(branchCmd),
	)
}
