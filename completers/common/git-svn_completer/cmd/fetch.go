package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download new revisions from SVN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().Bool("add-author-from", false, "Add author from")
	fetchCmd.Flags().BoolP("all", "all", false, "Fetch from all SVN remotes")
	fetchCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	fetchCmd.Flags().String("authors-prog", "", "Authors program")
	fetchCmd.Flags().Bool("follow-parent", false, "Follow parent")
	fetchCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	fetchCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	fetchCmd.Flags().String("include-paths", "", "Regex of paths to include")
	fetchCmd.Flags().Bool("localtime", false, "Store Git commit times in local time zone")
	fetchCmd.Flags().Int("log-window-size", 100, "Log window size")
	fetchCmd.Flags().Bool("no-checkout", false, "No checkout")
	fetchCmd.Flags().Bool("noMetadata", false, "Disable metadata")
	fetchCmd.Flags().BoolP("parent", "p", false, "Fetch only from the SVN parent of the current HEAD")
	fetchCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	fetchCmd.Flags().StringP("revision", "r", "", "Revision range")
	fetchCmd.Flags().Bool("use-log-author", false, "Use log author")
	fetchCmd.Flags().Bool("useSvmProps", false, "Use SVM properties")
	fetchCmd.Flags().Bool("useSvnsyncProps", false, "Use SVNSync properties")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).FlagCompletion(carapace.ActionMap{
		"authors-file":  carapace.ActionFiles(),
		"authors-prog":  carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
		"revision":      carapace.ActionValues(),
	})
}
