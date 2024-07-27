package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireMoveCmd = &cobra.Command{
	Use:   "livewire:move [--force] [--inline] [--] <name> <new-name>",
	Short: "Move a Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireMoveCmd).Standalone()

	livewireMoveCmd.Flags().Bool("force", false, "")
	livewireMoveCmd.Flags().Bool("inline", false, "")
	rootCmd.AddCommand(livewireMoveCmd)
}
