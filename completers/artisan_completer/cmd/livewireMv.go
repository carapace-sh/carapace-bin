package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireMvCmd = &cobra.Command{
	Use:   "livewire:mv [--inline] [--force] [--] <name> <new-name>",
	Short: "Move a Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireMvCmd).Standalone()

	livewireMvCmd.Flags().Bool("force", false, "")
	livewireMvCmd.Flags().Bool("inline", false, "")
	rootCmd.AddCommand(livewireMvCmd)
}
