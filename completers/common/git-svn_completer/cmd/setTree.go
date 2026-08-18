package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setTreeCmd = &cobra.Command{
	Use:   "set-tree",
	Short: "Set an SVN repository to a git tree-ish",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setTreeCmd).Standalone()

	setTreeCmd.Flags().Bool("add-author-from", false, "Add author from")
	setTreeCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	setTreeCmd.Flags().String("authors-prog", "", "Authors program")
	setTreeCmd.Flags().IntP("copy-similarity", "C", 0, "Copy similarity threshold")
	setTreeCmd.Flags().BoolP("edit", "e", false, "Edit commit message")
	setTreeCmd.Flags().BoolP("find-copies-harder", "l", false, "Find copies harder")
	setTreeCmd.Flags().Bool("follow-parent", false, "Follow parent")
	setTreeCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	setTreeCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	setTreeCmd.Flags().String("include-paths", "", "Regex of paths to include")
	setTreeCmd.Flags().Bool("localtime", false, "Store Git commit times in local time zone")
	setTreeCmd.Flags().Int("log-window-size", 100, "Log window size")
	setTreeCmd.Flags().Bool("no-checkout", false, "No checkout")
	setTreeCmd.Flags().Bool("noMetadata", false, "Disable metadata")
	setTreeCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	setTreeCmd.Flags().Bool("rmdir", false, "Remove empty directories")
	setTreeCmd.Flags().Bool("stdin", false, "Read a list of commits from stdin")
	setTreeCmd.Flags().Bool("use-log-author", false, "Use log author")
	setTreeCmd.Flags().Bool("useSvmProps", false, "Use SVM properties")
	setTreeCmd.Flags().Bool("useSvnsyncProps", false, "Use SVNSync properties")
	rootCmd.AddCommand(setTreeCmd)

	carapace.Gen(setTreeCmd).FlagCompletion(carapace.ActionMap{
		"authors-file":  carapace.ActionFiles(),
		"authors-prog":  carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
	})

	carapace.Gen(setTreeCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
