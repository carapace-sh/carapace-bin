package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a repo for tracking (requires URL argument)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringArrayP("branches", "b", []string{}, "Branches path")
	initCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	initCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	initCmd.Flags().String("include-paths", "", "Regex of paths to include")
	initCmd.Flags().Bool("minimize-url", false, "Minimize URL")
	initCmd.Flags().Bool("no-metadata", false, "Disable metadata")
	initCmd.Flags().BoolP("no-minimize-url", "m", false, "Do not minimize URL")
	initCmd.Flags().String("prefix", "", "Remote ref prefix")
	initCmd.Flags().String("rewrite-root", "", "Rewrite root URL")
	initCmd.Flags().String("rewrite-uuid", "", "Rewrite UUID")
	initCmd.Flags().Bool("shared", false, "Shared repository")
	initCmd.Flags().BoolP("stdlayout", "s", false, "Use standard layout")
	initCmd.Flags().StringArrayP("tags", "t", []string{}, "Tags path")
	initCmd.Flags().String("template", "", "Git template directory")
	initCmd.Flags().StringP("trunk", "T", "", "Trunk path")
	initCmd.Flags().Bool("use-svm-props", false, "Use SVM properties")
	initCmd.Flags().Bool("use-svnsync-props", false, "Use SVNSync properties")
	initCmd.Flags().String("username", "", "SVN username")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"branches":      carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
		"tags":          carapace.ActionFiles(),
		"template":      carapace.ActionFiles(),
		"trunk":         carapace.ActionFiles(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
