package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiInitCmd = &cobra.Command{
	Use:   "multi-init",
	Short: "Deprecated alias for 'git svn init -T<trunk> -b<branches> -t<tags>'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiInitCmd).Standalone()

	multiInitCmd.Flags().StringArrayP("branches", "b", []string{}, "Branches path")
	multiInitCmd.Flags().String("config-dir", "", "SVN configuration directory")
	multiInitCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	multiInitCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	multiInitCmd.Flags().String("include-paths", "", "Regex of paths to include")
	multiInitCmd.Flags().BoolP("minimize-url", "m", false, "Minimize URL")
	multiInitCmd.Flags().Bool("no-auth-cache", false, "Disable SVN authentication caching")
	multiInitCmd.Flags().Bool("no-metadata", false, "Disable metadata")
	multiInitCmd.Flags().Bool("no-minimize-url", false, "Do not minimize URL")
	multiInitCmd.Flags().String("prefix", "", "Remote ref prefix")
	multiInitCmd.Flags().String("rewrite-root", "", "Rewrite root URL")
	multiInitCmd.Flags().String("rewrite-uuid", "", "Rewrite UUID")
	multiInitCmd.Flags().String("shared", "", "Shared repository")
	multiInitCmd.Flags().BoolP("stdlayout", "s", false, "Use standard layout")
	multiInitCmd.Flags().StringArrayP("tags", "t", []string{}, "Tags path")
	multiInitCmd.Flags().String("template", "", "Git template directory")
	multiInitCmd.Flags().StringP("trunk", "T", "", "Trunk path")
	multiInitCmd.Flags().Bool("use-svm-props", false, "Use SVM properties")
	multiInitCmd.Flags().Bool("use-svnsync-props", false, "Use SVNSync properties")
	multiInitCmd.Flags().String("username", "", "SVN username")
	rootCmd.AddCommand(multiInitCmd)

	carapace.Gen(multiInitCmd).FlagCompletion(carapace.ActionMap{
		"branches":      carapace.ActionFiles(),
		"config-dir":    carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
		"tags":          carapace.ActionFiles(),
		"template":      carapace.ActionFiles(),
		"trunk":         carapace.ActionFiles(),
		"username":      carapace.ActionValues(),
	})

	carapace.Gen(multiInitCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
