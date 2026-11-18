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
	flagCacheCmd.Flags().String("arch", "", "Show cache file for the given CPU architecture. (Pass `all` to show cache files for all architectures.)")
	flagCacheCmd.Flags().String("bottle-tag", "", "Show the cache file used when pouring a bottle for the given tag.")
	flagCacheCmd.Flags().BoolP("build-from-source", "s", false, "Show the cache file used when building from source.")
	flagCacheCmd.Flags().Bool("cask", false, "Only show cache files for casks.")
	flagCacheCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagCacheCmd.Flags().Bool("force-bottle", false, "Show the cache file used when pouring a bottle.")
	flagCacheCmd.Flags().Bool("formula", false, "Only show cache files for formulae.")
	flagCacheCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagCacheCmd.Flags().String("os", "", "Show cache file for the given operating system. (Pass `all` to show cache files for all operating systems.)")
	flagCacheCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagCacheCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")

	carapace.Gen(flagCacheCmd).FlagCompletion(carapace.ActionMap{
		"arch":       carapace.ActionValues("all", "arm64", "x86_64"),
		"bottle-tag": carapace.ActionValues(),
		"os":         carapace.ActionValues("all", "linux", "macos"),
	})
}
