package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getversionCommand = &cobra.Command{
	Use:   "getversion <pkgver|dep...>",
	Short: "Prints version from patterns or pkgvers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getversionCommand).Standalone()

	rootCmd.AddCommand(getversionCommand)
}
