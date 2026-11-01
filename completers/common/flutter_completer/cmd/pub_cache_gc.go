package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_cache_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Prunes unused packages from the system cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_cache_gcCmd).Standalone()

	pub_cache_gcCmd.Flags().Bool("collect-recent", false, "Also delete recent files")
	pub_cache_gcCmd.Flags().Bool("dry-run", false, "Print list of files that would be deleted")
	pub_cache_gcCmd.Flags().BoolP("force", "f", false, "Prune cache without confirmation")
	pub_cache_gcCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_cacheCmd.AddCommand(pub_cache_gcCmd)
}
