package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-create [OPTIONS] -A <arch> -n <pkgver> -s \"<desc>\" destdir",
	Short: "XBPS utility to create binary packages",
	Long:  "https://man.voidlinux.org/xbps-create",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("architecture", "A", "", "Package architecture (e.g: noarch, i686, etc)")
	rootCmd.Flags().StringP("built-with", "B", "", "Package builder string (e.g: xbps-src-30)")
	rootCmd.Flags().StringP("conflicts", "C", "", "Conflicts (blank separated list, e.g: 'foo>=2.0 blah<=2.0')")
	rootCmd.Flags().StringP("changelog", "c", "", "Changelog URL")
	rootCmd.Flags().StringP("dependencies", "D", "", "Dependencies (blank separated list, e.g: 'foo>=1.0_1 blah<2.1')")
	rootCmd.Flags().StringP("config-files", "F", "", "Configuration files (blank separated list)")
	rootCmd.Flags().StringP("homepage", "H", "", "Homepage")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().StringP("license", "l", "", "License")
	rootCmd.Flags().StringP("mutable-files", "M", "", "Mutable files list (blank separated list)")
	rootCmd.Flags().StringP("maintainer", "m", "", "Maintainer")
	rootCmd.Flags().StringP("pkgver", "n", "", "Package name/version tuple (e.g: foo-1.0_1)")
	rootCmd.Flags().StringP("provides", "P", "", "Provides (blank separated list, e.g: 'foo-9999 blah-1.0')")
	rootCmd.Flags().BoolP("preserve", "p", false, "Enable package preserve boolean")
	rootCmd.Flags().BoolP("quiet", "q", false, "Work silently")
	rootCmd.Flags().StringP("replaces", "R", "", "Replaces (blank separated list, e.g: 'foo>=1.0 blah<2.0')")
	rootCmd.Flags().StringP("reverts", "r", "", "Reverts (blank separated list, e.g: '1.0_1 2.0_3')")
	rootCmd.Flags().StringP("long-desc", "S", "", "Long description (80 cols per line)")
	rootCmd.Flags().StringP("desc", "s", "", "Short description (max 80 characters)")
	rootCmd.Flags().StringP("tags", "t", "", "A list of tags/categories (blank separated list)")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")
	rootCmd.Flags().String("alternatives", "", "List of available alternatives this pkg provides")
	rootCmd.Flags().String("build-options", "", "A string with the used build options")
	rootCmd.Flags().String("compression", "zstd", "Compression format: none, gzip, bzip2, lz4, xz, zstd (default)")
	rootCmd.Flags().String("shlib-provides", "", "List of provided shared libraries")
	rootCmd.Flags().String("shlib-requires", "", "List of required shared libraries")

	rootCmd.MarkFlagRequired("architecture")
	rootCmd.MarkFlagRequired("pkgver")
	rootCmd.MarkFlagRequired("desc")

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionDirectories())
}
