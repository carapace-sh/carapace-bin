package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Output SVG graphic showing service initialization",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plotCmd).Standalone()

	rootCmd.AddCommand(plotCmd)
}
