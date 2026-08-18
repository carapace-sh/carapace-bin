package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Initialize and fetch revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().Bool("add-author-from", false, "Add author from")
	cloneCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	cloneCmd.Flags().String("authors-prog", "", "Authors program")
	cloneCmd.Flags().StringArrayP("branches", "b", []string{}, "Branches path")
	cloneCmd.Flags().Bool("follow-parent", false, "Follow parent")
	cloneCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	cloneCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	cloneCmd.Flags().String("include-paths", "", "Regex of paths to include")
	cloneCmd.Flags().Bool("localtime", false, "Store Git commit times in local time zone")
	cloneCmd.Flags().Int("log-window-size", 100, "Log window size")
	cloneCmd.Flags().Bool("minimize-url", false, "Minimize URL")
	cloneCmd.Flags().Bool("no-checkout", false, "No checkout")
	cloneCmd.Flags().Bool("no-metadata", false, "Disable metadata")
	cloneCmd.Flags().Bool("no-minimize-url", false, "Do not minimize URL")
	cloneCmd.Flags().String("placeholder-filename", "", "Placeholder filename")
	cloneCmd.Flags().String("prefix", "", "Remote ref prefix")
	cloneCmd.Flags().Bool("preserve-empty-dirs", false, "Preserve empty directories")
	cloneCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	cloneCmd.Flags().StringP("revision", "r", "", "Revision range")
	cloneCmd.Flags().Bool("shared", false, "Shared repository")
	cloneCmd.Flags().BoolP("stdlayout", "s", false, "Use standard layout")
	cloneCmd.Flags().StringArrayP("tags", "t", []string{}, "Tags path")
	cloneCmd.Flags().String("template", "", "Git template directory")
	cloneCmd.Flags().StringP("trunk", "T", "", "Trunk path")
	cloneCmd.Flags().Bool("use-log-author", false, "Use log author")
	cloneCmd.Flags().Bool("use-svm-props", false, "Use SVM properties")
	cloneCmd.Flags().Bool("use-svnsync-props", false, "Use SVNSync properties")
	cloneCmd.Flags().Bool("useSvmProps", false, "Use SVM properties")
	cloneCmd.Flags().Bool("useSvnsyncProps", false, "Use SVNSync properties")
	cloneCmd.Flags().String("username", "", "SVN username")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"authors-file":         carapace.ActionFiles(),
		"authors-prog":         carapace.ActionFiles(),
		"branches":             carapace.ActionFiles(),
		"ignore-paths":         carapace.ActionValues(),
		"ignore-refs":          carapace.ActionValues(),
		"include-paths":        carapace.ActionValues(),
		"placeholder-filename": carapace.ActionFiles(),
		"revision":             carapace.ActionValues(),
		"tags":                 carapace.ActionFiles(),
		"template":             carapace.ActionFiles(),
		"trunk":                carapace.ActionFiles(),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)
}
