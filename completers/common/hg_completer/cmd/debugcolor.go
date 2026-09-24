package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugcolorCmd = &cobra.Command{
	Use:    "debugcolor",
	Short:  "hg debugcolor",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugcolorCmd).Standalone()

	debugcolorCmd.Flags().Bool("style", false, "show all configured styles")
	rootCmd.AddCommand(debugcolorCmd)
}
