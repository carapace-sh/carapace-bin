package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireDeleteCmd = &cobra.Command{
	Use:   "livewire:delete [--inline] [--force] [--test] [--] <name>",
	Short: "Delete a Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireDeleteCmd).Standalone()

	livewireDeleteCmd.Flags().Bool("force", false, "")
	livewireDeleteCmd.Flags().Bool("inline", false, "")
	livewireDeleteCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(livewireDeleteCmd)
}
