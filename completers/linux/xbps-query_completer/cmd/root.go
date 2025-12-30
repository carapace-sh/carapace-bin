package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-query [OPTIONS] MODE [ARGUMENTS]",
	Short: "XBPS utility to query for package and repository information",
	Long:  "https://man.voidlinux.org/xbps-query",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().StringP("cachedir", "c", "", "Path to cachedir")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("ignore-conf-repos", "i", false, "Ignore repositories defined in xbps.d")
	rootCmd.Flags().BoolP("memory-sync", "M", false, "Remote repository data is fetched and stored in memory, ignoring on-disk repodata archives")
	rootCmd.Flags().StringP("property", "p", "", "Show properties for PKGNAME")
	rootCmd.Flags().BoolP("R", "R", false, "Enable repository mode. This mode explicitly looks for packages in repositories")
	rootCmd.Flags().StringArray("repository", nil, "Add repository to the top of the list. This option can be specified multiple times")
	rootCmd.Flags().Bool("staging", false, "Enable use of staged packages")
	rootCmd.Flags().Bool("regex", false, "Use Extended Regular Expressions to match")
	rootCmd.Flags().Bool("fulldeptree", false, "Full dependency tree for -x/--deps")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")
	rootCmd.Flags().BoolP("list-pkgs", "l", false, "List installed packages")
	rootCmd.Flags().BoolP("list-repos", "L", false, "List registered repositories")
	rootCmd.Flags().BoolP("list-hold-pkgs", "H", false, "List packages on hold state")
	rootCmd.Flags().Bool("list-repolock-pkgs", false, "List repolocked packages")
	rootCmd.Flags().BoolP("list-manual-pkgs", "m", false, "List packages installed explicitly")
	rootCmd.Flags().BoolP("list-orphans", "O", false, "List package orphans")
	rootCmd.Flags().StringP("ownedby", "o", "", "Search for package files by matching STRING or REGEX")
	rootCmd.Flags().StringP("show", "S", "", "Show information for PKG [default mode]")
	rootCmd.Flags().StringP("search", "s", "", "Search for packages by matching PKG, STRING or REGEX")
	rootCmd.Flags().String("cat", "", "Print FILE from PKG binpkg to stdout")
	rootCmd.Flags().StringP("files", "f", "", "Show package files for PKG")
	rootCmd.Flags().StringP("deps", "x", "", "Show dependencies for PKG")
	rootCmd.Flags().StringP("revdeps", "X", "", "Show reverse dependencies for PKG")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionDirectories(),
		"cachedir": carapace.ActionDirectories(),
		"rootdir":  carapace.ActionDirectories(),
		"ownedby":  carapace.ActionFiles(),
		"show":     xbps.ActionPackageSearch(),
		"search":   xbps.ActionPackageSearch(),
		"files":    xbps.ActionPackageSearch(),
		"deps":     xbps.ActionPackageSearch(),
		"revdeps":  xbps.ActionPackageSearch(),
	})
}
