package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packageDiscoverCmd = &cobra.Command{
	Use:   "package:discover",
	Short: "Rebuild the cached package manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packageDiscoverCmd).Standalone()

	rootCmd.AddCommand(packageDiscoverCmd)
}
