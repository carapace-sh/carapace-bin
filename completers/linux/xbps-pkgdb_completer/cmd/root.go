package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-pkgdb [OPTIONS] [PKGNAME...]",
	Short: "XBPS utility to report/fix issues and modify the package database (pkgdb)",
	Long:  "https://man.voidlinux.org/xbps-pkgdb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Process all packages")
	rootCmd.Flags().StringSlice("checks", nil, "Choose checks to run")
	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().StringP("mode", "m", "", "Change PKGNAME to this mode")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().BoolP("update", "u", false, "Update pkgdb to the latest format")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"checks":  carapace.ActionValues("files", "dependencies", "alternatives", "pkgdb"),
		"config":  carapace.ActionDirectories(),
		"mode":    carapace.ActionValues("auto", "manual", "hold", "unhold", "repolock", "repounlock"),
		"rootdir": carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(xbps.ActionPackageSearch())
}
