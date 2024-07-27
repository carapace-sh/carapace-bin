package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireCopyCmd = &cobra.Command{
	Use:   "livewire:copy [--inline] [--force] [--test] [--] <name> <new-name>",
	Short: "Copy a Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireCopyCmd).Standalone()

	livewireCopyCmd.Flags().Bool("force", false, "")
	livewireCopyCmd.Flags().Bool("inline", false, "")
	livewireCopyCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(livewireCopyCmd)
}
