package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-remove [OPTIONS] [PKGNAME...]",
	Short: "XBPS utility to remove packages",
	Long:  "https://man.voidlinux.org/xbps-remove",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("cachedir", "c", "", "Path to cachedir")
	rootCmd.Flags().BoolP("clean-cache", "O", false, "Remove outdated packages from the cache. If specified twice, also remove uninstalled packages")
	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("dry-run", "n", false, "Dry-run mode")
	rootCmd.Flags().BoolP("force", "f", false, "Force package re-installation. If specified twice, all files will be overwritten.")
	rootCmd.Flags().BoolP("force-revdeps", "F", false, "Force package removal even with revdeps or unresolved shared libraries")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("recursive", "R", false, "Recursively remove dependencies")
	rootCmd.Flags().BoolP("remove-orphans", "o", false, "Remove package orphans")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")
	rootCmd.Flags().BoolP("yes", "y", false, "Assume yes to all questions")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cachedir": carapace.ActionDirectories(),
		"config":   carapace.ActionDirectories(),
		"rootdir":  carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(xbps.ActionPackages().FilterArgs())
}
