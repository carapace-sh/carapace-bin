package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Install a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_cache_addCmd).Standalone()

	pub_cache_addCmd.Flags().Bool("all", false, "Install all matching versions.")
	pub_cache_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_cache_addCmd.Flags().StringP("version", "v", "", "Version constraint.")
	pub_cacheCmd.AddCommand(pub_cache_addCmd)
}
