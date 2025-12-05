package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pamac"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display package details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().BoolP("aur", "a", false, "also search in AUR")
	infoCmd.Flags().Bool("no-aur", false, "do not search in AUR")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		pamac.ActionPackageSearch().FilterArgs(), // TODO aur flag
	)
}
