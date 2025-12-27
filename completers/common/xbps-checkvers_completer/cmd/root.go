package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-checkvers [OPTIONS] [FILES...]",
	Short: "XBPS utility to check for outdated packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().StringP("distdir", "D", "", "Set (or override) the path to void-packages (defaults to ~/void-packages)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("removed", "e", false, "List packages present in repos, but not in distdir")
	rootCmd.Flags().BoolP("installed", "I", false, "Check for outdated packages in rootdir, rather than in the XBPS repositories")
	rootCmd.Flags().BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	rootCmd.Flags().BoolP("manual", "m", false, "Only process listed files")
	rootCmd.Flags().StringArrayP("repository", "R", nil, "Append repository to the head of repository list")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().BoolP("show-all", "s", false, "List all packages, in the format 'pkgname repover srcver'")
	rootCmd.Flags().Bool("staging", false, "Enable use of staged packages")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionDirectories(),
		"distdir": carapace.ActionDirectories(),
		"rootdir": carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
