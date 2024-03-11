package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagCacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Display Homebrew's download cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagCacheCmd).Standalone()

	flagCacheCmd.Flags().Bool("HEAD", false, "Show the cache file used when building from HEAD.")
	flagCacheCmd.Flags().Bool("arch", false, "Show cache file for the given CPU architecture. (Pass `all` to show cache files for all architectures.)")
	flagCacheCmd.Flags().Bool("bottle-tag", false, "Show the cache file used when pouring a bottle for the given tag.")
	flagCacheCmd.Flags().Bool("build-from-source", false, "Show the cache file used when building from source.")
	flagCacheCmd.Flags().Bool("cask", false, "Only show cache files for casks.")
	flagCacheCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagCacheCmd.Flags().Bool("force-bottle", false, "Show the cache file used when pouring a bottle.")
	flagCacheCmd.Flags().Bool("formula", false, "Only show cache files for formulae.")
	flagCacheCmd.Flags().Bool("help", false, "Show this message.")
	flagCacheCmd.Flags().Bool("os", false, "Show cache file for the given operating system. (Pass `all` to show cache files for all operating systems.)")
	flagCacheCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagCacheCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
