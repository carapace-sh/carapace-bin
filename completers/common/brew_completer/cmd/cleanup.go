package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:     "cleanup",
	Short:   "Remove stale lock files and outdated downloads for all formulae and casks, and remove old versions of installed formulae",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanupCmd).Standalone()

	cleanupCmd.Flags().Bool("debug", false, "Display any debugging information.")
	cleanupCmd.Flags().Bool("dry-run", false, "Show what would be removed, but do not actually remove anything.")
	cleanupCmd.Flags().Bool("help", false, "Show this message.")
	cleanupCmd.Flags().String("prune", "", "Remove all cache files older than specified <days>. If you want to remove everything, use `--prune=all`.")
	cleanupCmd.Flags().Bool("prune-prefix", false, "Only prune the symlinks and directories from the prefix and remove no other files.")
	cleanupCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	cleanupCmd.Flags().BoolS("s", "s", false, "Scrub the cache, including downloads for even the latest versions. Note that downloads for any installed formulae or casks will still not be deleted. If you want to delete those too: `rm -rf \"$(brew --cache)\"`")
	cleanupCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(cleanupCmd)

	carapace.Gen(cleanupCmd).FlagCompletion(carapace.ActionMap{
		"prune": carapace.ActionValues("all"),
	})
}
