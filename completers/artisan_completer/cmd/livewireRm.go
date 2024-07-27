package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireRmCmd = &cobra.Command{
	Use:   "livewire:rm [--inline] [--force] [--test] [--] <name>",
	Short: "Delete a Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireRmCmd).Standalone()

	livewireRmCmd.Flags().Bool("force", false, "")
	livewireRmCmd.Flags().Bool("inline", false, "")
	livewireRmCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(livewireRmCmd)
}
