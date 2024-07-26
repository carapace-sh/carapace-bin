package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireTouchCmd = &cobra.Command{
	Use:   "livewire:touch [--force] [--inline] [--test] [--pest] [--stub [STUB]] [--] <name>",
	Short: "Create a new Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireTouchCmd).Standalone()

	livewireTouchCmd.Flags().Bool("force", false, "")
	livewireTouchCmd.Flags().Bool("inline", false, "")
	livewireTouchCmd.Flags().Bool("pest", false, "")
	livewireTouchCmd.Flags().String("stub", "", "")
	livewireTouchCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(livewireTouchCmd)
}
