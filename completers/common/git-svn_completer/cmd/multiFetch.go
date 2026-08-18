package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiFetchCmd = &cobra.Command{
	Use:   "multi-fetch",
	Short: "Deprecated alias for 'git svn fetch --all'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiFetchCmd).Standalone()

	multiFetchCmd.Flags().Bool("add-author-from", false, "Add author from")
	multiFetchCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	multiFetchCmd.Flags().String("authors-prog", "", "Authors program")
	multiFetchCmd.Flags().String("config-dir", "", "SVN configuration directory")
	multiFetchCmd.Flags().Bool("follow-parent", false, "Follow parent")
	multiFetchCmd.Flags().String("ignore-paths", "", "Regex of paths to ignore")
	multiFetchCmd.Flags().String("ignore-refs", "", "Regex of SVN refs to ignore")
	multiFetchCmd.Flags().String("include-paths", "", "Regex of paths to include")
	multiFetchCmd.Flags().Bool("localtime", false, "Store Git commit times in local time zone")
	multiFetchCmd.Flags().Int("log-window-size", 100, "Log window size")
	multiFetchCmd.Flags().Bool("no-auth-cache", false, "Disable SVN authentication caching")
	multiFetchCmd.Flags().Bool("no-checkout", false, "No checkout")
	multiFetchCmd.Flags().Bool("noMetadata", false, "Disable metadata")
	multiFetchCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	multiFetchCmd.Flags().Int("repack", 0, "Repack interval")
	multiFetchCmd.Flags().String("repack-flags", "", "Flags to pass to repack")
	multiFetchCmd.Flags().StringP("revision", "r", "", "Revision range")
	multiFetchCmd.Flags().Bool("use-log-author", false, "Use log author")
	multiFetchCmd.Flags().Bool("useSvmProps", false, "Use SVM properties")
	multiFetchCmd.Flags().Bool("useSvnsyncProps", false, "Use SVNSync properties")
	multiFetchCmd.Flags().String("username", "", "SVN username")
	rootCmd.AddCommand(multiFetchCmd)

	carapace.Gen(multiFetchCmd).FlagCompletion(carapace.ActionMap{
		"authors-file":  carapace.ActionFiles(),
		"authors-prog":  carapace.ActionFiles(),
		"config-dir":    carapace.ActionFiles(),
		"ignore-paths":  carapace.ActionValues(),
		"ignore-refs":   carapace.ActionValues(),
		"include-paths": carapace.ActionValues(),
		"repack-flags":  carapace.ActionValues(),
		"revision":      carapace.ActionValues(),
		"username":      carapace.ActionValues(),
	})
}
