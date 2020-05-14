package cmd

import (
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use: "pull",
	Short: "Fetch from and integrate with another repository or a local branch",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	pullCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	pullCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	pullCmd.Flags().BoolP("append", "a", false, "append to .git/FETCH_HEAD instead of overwriting")
	pullCmd.Flags().Bool("all", false, "fetch from all remotes")
	pullCmd.Flags().Bool("allow-unrelated-histories", false, "allow merging unrelated histories")
	pullCmd.Flags().Bool("autostash", false, "automatically stash/stash pop before and after rebase")
	pullCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	pullCmd.Flags().Bool("commit", false, "perform a commit if the merge succeeds (default)")
	pullCmd.Flags().String("depth", "", "deepen history of shallow clone")
	pullCmd.Flags().Bool("dry-run", false, "dry run")
	pullCmd.Flags().Bool("edit", false, "edit message before committing")
	pullCmd.Flags().Bool("ff", false, "allow fast-forward")
	pullCmd.Flags().Bool("ff-only", false, "abort if fast-forward is not possible")
	pullCmd.Flags().BoolP("force", "f", false, "force overwrite of local branch")
	pullCmd.Flags().StringP("jobs", "j", "", "number of submodules pulled in parallel")
	pullCmd.Flags().BoolP("keep", "k", false, "keep downloaded pack")
	pullCmd.Flags().String("log", "", "add (at most <n>) entries from shortlog to merge commit message")
	pullCmd.Flags().BoolP("n", "n", false, "do not show a diffstat at the end of the merge")
	pullCmd.Flags().BoolP("prune", "p", false, "prune remote-tracking branches no longer on remote")
	pullCmd.Flags().Bool("progress", false, "force progress reporting")
	pullCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	pullCmd.Flags().String("recurse-submodules", "", "control for recursive fetching of submodules")
	pullCmd.Flags().String("refmap", "", "specify fetch refmap")
	pullCmd.Flags().StringP("rebase", "r", "", "incorporate changes by rebasing rather than merging")
	pullCmd.Flags().Bool("set-upstream", false, "set upstream for git pull/fetch")
	pullCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	pullCmd.Flags().Bool("show-forced-updates", false, "check for forced-updates on all updated branches")
	pullCmd.Flags().String("signoff", "", "add Signed-off-by:")
	pullCmd.Flags().Bool("squash", false, "create a single commit instead of doing a merge")
	pullCmd.Flags().BoolP("strategy", "s", false, "<strategy>    merge strategy to use")
	pullCmd.Flags().Bool("stat", false, "show a diffstat at the end of the merge")
	pullCmd.Flags().BoolP("tags", "t", false, "fetch all tags and associated objects")
	pullCmd.Flags().Bool("unshallow", false, "convert to a complete repository")
	pullCmd.Flags().Bool("update-shallow", false, "accept refs that update .git/shallow")
	pullCmd.Flags().String("upload-pack", "", "path to upload pack on remote end")
	pullCmd.Flags().Bool("verify-signatures", false, "verify that the named commit has a valid GPG signature")
	pullCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	pullCmd.Flags().BoolP("strategy-option", "X", false, "<option=value>    option for selected merge strategy")
    rootCmd.AddCommand(pullCmd)
}
