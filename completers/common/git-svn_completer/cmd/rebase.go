package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Fetch and rebase your working directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().Bool("add-author-from", false, "Add author from")
	rebaseCmd.Flags().Bool("all", false, "Fetch from all SVN remotes")
	rebaseCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	rebaseCmd.Flags().String("authors-prog", "", "Authors program")
	rebaseCmd.Flags().String("config-dir", "", "SVN configuration directory")
	rebaseCmd.Flags().BoolP("dry-run", "n", false, "Dry run")
	rebaseCmd.Flags().Bool("follow-parent", false, "Follow parent")
	rebaseCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	rebaseCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	rebaseCmd.Flags().String("include-paths", "", "Regex of paths to include")
	rebaseCmd.Flags().BoolP("local", "l", false, "Do not fetch remotely; only run git rebase")
	rebaseCmd.Flags().Bool("localtime", false, "Store Git commit times in local time zone")
	rebaseCmd.Flags().Int("log-window-size", 100, "Log window size")
	rebaseCmd.Flags().BoolP("merge", "m", false, "Merge")
	rebaseCmd.Flags().Bool("no-auth-cache", false, "Disable SVN authentication caching")
	rebaseCmd.Flags().Bool("no-checkout", false, "No checkout")
	rebaseCmd.Flags().Bool("noMetadata", false, "Disable metadata")
	rebaseCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	rebaseCmd.Flags().BoolP("rebase-merges", "p", false, "Rebase merges")
	rebaseCmd.Flags().Int("repack", 0, "Repack interval")
	rebaseCmd.Flags().String("repack-flags", "", "Flags to pass to repack")
	rebaseCmd.Flags().StringP("strategy", "s", "", "Rebase strategy")
	rebaseCmd.Flags().Bool("use-log-author", false, "Use log author")
	rebaseCmd.Flags().Bool("useSvmProps", false, "Use SVM properties")
	rebaseCmd.Flags().Bool("useSvnsyncProps", false, "Use SVNSync properties")
	rebaseCmd.Flags().String("username", "", "SVN username")
	rebaseCmd.Flags().BoolP("verbose", "v", false, "Verbose")
	rootCmd.AddCommand(rebaseCmd)

	carapace.Gen(rebaseCmd).FlagCompletion(carapace.ActionMap{
		"authors-file":  carapace.ActionFiles(),
		"authors-prog":  carapace.ActionFiles(),
		"config-dir":    carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
		"repack-flags":  carapace.ActionValues(),
		"strategy":      carapace.ActionValues("recursive", "resolve", "ours", "theirs", "subtree"),
		"username":      carapace.ActionValues(),
	})
}
