package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkgmatchCommand = &cobra.Command{
	Use:   "getversion <pkgver> <dep>",
	Short: "Match pkgver against dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkgmatchCommand).Standalone()

	rootCmd.AddCommand(pkgmatchCommand)
}
