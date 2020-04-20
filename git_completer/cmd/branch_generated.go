package cmd

import (
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use: "branch",
	Short: "List, create, or delete branches",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	branchCmd.Flags().BoolP("all", "a", false, "list both remote-tracking and local branches")
	branchCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	branchCmd.Flags().BoolP("C", "C", false, "copy a branch, even if target exists")
	branchCmd.Flags().BoolP("copy", "c", false, "copy a branch and its reflog")
	branchCmd.Flags().String("color", "", "use colored output")
	branchCmd.Flags().String("column", "", "list branches in columns")
	branchCmd.Flags().String("contains", "", "print only branches that contain the commit")
	branchCmd.Flags().Bool("create-reflog", false, "create the branch's reflog")
	branchCmd.Flags().BoolP("D", "D", false, "delete branch (even if not merged)")
	branchCmd.Flags().BoolP("delete", "d", false, "delete fully merged branch")
	branchCmd.Flags().Bool("edit-description", false, "edit the description for the branch")
	branchCmd.Flags().BoolP("force", "f", false, "force creation, move/rename, deletion")
	branchCmd.Flags().String("format", "", "format to use for the output")
	branchCmd.Flags().BoolP("ignore-case", "i", false, "sorting and filtering are case insensitive")
	branchCmd.Flags().BoolP("list", "l", false, "list branch names")
	branchCmd.Flags().String("merged", "", "print only branches that are merged")
	branchCmd.Flags().BoolP("move", "m", false, "move/rename a branch and its reflog")
	branchCmd.Flags().BoolP("M", "M", false, "move/rename a branch, even if target exists")
	branchCmd.Flags().String("no-contains", "", "print only branches that don't contain the commit")
	branchCmd.Flags().String("no-merged", "", "print only branches that are not merged")
	branchCmd.Flags().String("points-at", "", "print only branches of the object")
	branchCmd.Flags().BoolP("quiet", "q", false, "suppress informational messages")
	branchCmd.Flags().BoolP("remotes", "r", false, "act on remote-tracking branches")
	branchCmd.Flags().Bool("show-current", false, "show current branch name")
	branchCmd.Flags().String("sort", "", "field name to sort on")
	branchCmd.Flags().BoolP("track", "t", false, "set up tracking mode (see git-pull(1))")
	branchCmd.Flags().Bool("unset-upstream", false, "unset the upstream info")
	branchCmd.Flags().BoolP("set-upstream-to", "u", false, "<upstream>    change the upstream info")
	branchCmd.Flags().BoolP("verbose", "v", false, "show hash and subject, give twice for upstream branch")
    rootCmd.AddCommand(branchCmd)
}
