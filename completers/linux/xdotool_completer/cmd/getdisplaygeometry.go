package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getdisplaygeometryCmd = &cobra.Command{
	Use:   "getdisplaygeometry",
	Short: "Get the display geometry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getdisplaygeometryCmd).Standalone()

	rootCmd.AddCommand(getdisplaygeometryCmd)
}
