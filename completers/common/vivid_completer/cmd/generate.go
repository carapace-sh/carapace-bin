package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vivid"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a LS_COLORS expression",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).PositionalCompletion(
		vivid.ActionThemes(),
	)
}
