package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-install [OPTIONS] [PKGNAME...]",
	Short: "XBPS utility to (re)install and update packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("automatic", "A", false, "Set automatic installation mode")
	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().StringP("cachedir", "c", "", "Path to cachedir")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("download-only", "D", false, "Download packages and check integrity, nothing else")
	rootCmd.Flags().BoolSliceP("force", "f", nil, "Force package re-installation. If specified twice, all files will be overwritten.")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	rootCmd.Flags().BoolP("ignore-file-conflicts", "I", false, "Ignore detected file conflicts")
	rootCmd.Flags().BoolP("unpack-only", "U", false, "Unpack packages in transaction, do not configure them")
	rootCmd.Flags().BoolP("memory-sync", "M", false, "Remote repository data is fetched and stored in memory, ignoring on-disk repodata archives")
	rootCmd.Flags().BoolP("dry-run", "n", false, "Dry-run mode")
	rootCmd.Flags().StringArrayP("repository", "R", nil, "Add repository to the top of the list. This option can be specified multiple times")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().Bool("reproducible", false, "Enable reproducible mode in pkgdb")
	rootCmd.Flags().Bool("staging", false, "Enable use of staged packages")
	rootCmd.Flags().BoolP("sync", "S", false, "Sync remote repository index")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.Flags().BoolP("yes", "y", false, "Assume yes to all questions")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"cachedir": carapace.ActionDirectories(),
		"rootdir":  carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(xbps.ActionPackageSearch())
}
