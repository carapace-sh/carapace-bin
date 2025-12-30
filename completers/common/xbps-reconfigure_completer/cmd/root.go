package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-reconfigure [OPTIONS] [PKGNAME...]",
	Short: "XBPS utility to configure installed packages",
	Long:  "https://man.voidlinux.org/xbps-reconfigure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Process all packages")
	rootCmd.Flags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.Flags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.Flags().BoolP("force", "f", false, "Force package re-installation. If specified twice, all files will be overwritten.")
	rootCmd.Flags().Bool("fulldeptree", false, "Full dependency tree for -x/--deps")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().StringP("ignore", "i", "", "Ignore PKG with -a/--all")
	rootCmd.Flags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.Flags().BoolP("deps", "-x", false, "Also process dependencies for each package")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionDirectories(),
		"rootdir": carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(xbps.ActionPackages().FilterArgs())
}
