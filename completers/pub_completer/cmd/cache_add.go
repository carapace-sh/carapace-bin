package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Install a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_addCmd).Standalone()

	cache_addCmd.Flags().Bool("all", false, "Install all matching versions.")
	cache_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	cache_addCmd.Flags().BoolP("version", "v", false, "Version constraint.")
	cacheCmd.AddCommand(cache_addCmd)

	// TODO package completion
}
