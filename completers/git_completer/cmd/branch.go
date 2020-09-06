package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/git_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "List, create, or delete branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().BoolS("D", "D", false, "Shortcut for --delete --force.")
	branchCmd.Flags().BoolS("M", "M", false, "Shortcut for --move --force.")
	branchCmd.Flags().BoolS("C", "C", false, "Shortcut for --copy --force.")
	branchCmd.Flags().String("abbrev", "", "Alter the sha1’s minimum display length in the output listing.")
	branchCmd.Flags().BoolP("all", "a", false, "")
	branchCmd.Flags().String("color", "", "Color branches to highlight current, local, and remote-tracking branches.")
	branchCmd.Flags().String("contains", "", "Only list branches which contain the specified commit (HEAD if not specified).")
	branchCmd.Flags().BoolP("copy", "c", false, "Copy a branch and the corresponding reflog.")
	branchCmd.Flags().Bool("create-reflog", false, "Create the branch’s reflog.")
	branchCmd.Flags().BoolP("delete", "d", false, "Delete a branch.")
	branchCmd.Flags().Bool("edit-description", false, "Open an editor and edit the text to explain what the branch is for.")
	branchCmd.Flags().BoolP("force", "f", false, "Reset <branchname> to <startpoint>, even if <branchname> exists already.")
	branchCmd.Flags().String("format", "", "A string that interpolates %(fieldname) from a branch ref being shown and the object it points at.")
	branchCmd.Flags().BoolP("ignore-case", "i", false, "Sorting and filtering branches are case insensitive.")
	branchCmd.Flags().BoolP("list", "l", false, "List branches.")
	branchCmd.Flags().String("merged", "", "Only list branches whose tips are reachable from the specified commit (HEAD if not specified).")
	branchCmd.Flags().BoolP("move", "m", false, "Move/rename a branch and the corresponding reflog.")
	branchCmd.Flags().Bool("no-abbrev", false, "Display the full sha1s in the output listing rather than abbreviating them.")
	branchCmd.Flags().Bool("no-color", false, "Turn off branch colors, even when the configuration file gives the default to color output.")
	branchCmd.Flags().String("no-contains", "", "Only list branches which don’t contain the specified commit (HEAD if not specified).")
	branchCmd.Flags().String("no-merged", "", "Only list branches whose tips are not reachable from the specified commit (HEAD if not specified).")
	branchCmd.Flags().Bool("no-track", false, "Do not set up \"upstream\" configuration, even if the branch.autoSetupMerge configuration variable is true.")
	branchCmd.Flags().String("points-at", "", "Only list branches of the given object.")
	branchCmd.Flags().BoolP("quiet", "q", false, "Be more quiet when creating or deleting a branch, suppressing non-error messages.")
	branchCmd.Flags().BoolP("remotes", "r", false, "List or delete (if used with -d) the remote-tracking branches.")
	branchCmd.Flags().String("set-upstream-to", "", "")
	branchCmd.Flags().Bool("show-current", false, "Print the name of the current branch.")
	branchCmd.Flags().String("sort", "", "Sort based on the key given.")
	branchCmd.Flags().BoolP("track", "t", false, "When creating a new branch, set up branch.<name>.remote and branch.<name>.merge configuration entries to mark the start-point branch as \"upstream\" from the new branch.")
	branchCmd.Flags().Bool("unset-upstream", false, "Remove the upstream information for <branchname>.")
	rootCmd.AddCommand(branchCmd)

	carapace.Gen(branchCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "auto", "never"),
		// TODO contains - complete recent commits/branches
		// TODO merged - complete recent commits
		// TODO no-contains - complete recent commits
		// TODO no-merged - complete recent commits
		// TODO points-at
		"set-upstream-to": action.ActionRemoteBranches(),
		"sort":            ActionSortFields(),
	})

}
func ActionSortFields() carapace.Action {
	return carapace.ActionValuesDescribed(
		"authordate", "the date component of the author header-field",
		"authoremail", "the email component of the author header-field",
		"authorname", "the name component of the author header-field",
		"author", "the author header-field",
		"body", "the body of the message",
		"committerdate", "the date component of the committer header-field",
		"committeremail", "the email component of the committer header-field",
		"committername", "the name component of the committer header-field",
		"committer", "the committer header-field",
		"contents", "complete message",
		"creatordate", "the date component of the creator header-field",
		"creator", "the creator header-field",
		"deltabase", "object name of the delta base of the object",
		"HEAD", "* if HEAD matches ref or space otherwise",
		"numparent", "number of parent objects",
		"objectname", "object name (SHA-1)",
		"objectsize", "the size of the object",
		"object", "the object header-field",
		"objecttype", "the type of the object",
		"parent", "the parent header-field",
		"push", "name of a local ref which represents the @{push} location for the displayed ref",
		"refname", "name of the ref",
		"subject", "the subject of the message",
		"symref", "the ref which the given symbolic ref refers to",
		"taggerdate", "the date component of the tagger header-field",
		"taggeremail", "the email component of the tagger header-field",
		"taggername", "the name component of the tagger header-field",
		"tagger", "the tagger header-field",
		"tag", "the tag header-field",
		"trailers", "structured information in commit messages",
		"tree", "the tree header-field",
		"type", "the type header-field",
		"upstream", "name of a local ref which can be considered “upstream” from the displayed ref",
		"version:refname", "sort by versions",
	)
}
