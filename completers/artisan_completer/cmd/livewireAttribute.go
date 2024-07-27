package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireAttributeCmd = &cobra.Command{
	Use:   "livewire:attribute [--force] [--] <name>",
	Short: "Create a new Livewire attribute class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireAttributeCmd).Standalone()

	livewireAttributeCmd.Flags().Bool("force", false, "")
	rootCmd.AddCommand(livewireAttributeCmd)
}
