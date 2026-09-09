package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_purgeCacheCmd = &cobra.Command{
	Use:   "purge-cache",
	Short: "Purge the global repository cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_purgeCacheCmd).Standalone()
	package_purgeCacheCmd.Flags().SetInterspersed(false)

	package_purgeCacheCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_purgeCacheCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_purgeCacheCmd)
}
