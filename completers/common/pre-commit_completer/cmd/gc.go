package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Clean unused cached repos.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gcCmd).Standalone()

	gcCmd.Flags().String("color", "", "Whether to use color in output.")
	rootCmd.AddCommand(gcCmd)
}
