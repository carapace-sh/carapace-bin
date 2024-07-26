package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireCpCmd = &cobra.Command{
	Use:   "livewire:cp [--inline] [--force] [--test] [--] <name> <new-name>",
	Short: "Copy a Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireCpCmd).Standalone()

	livewireCpCmd.Flags().Bool("force", false, "")
	livewireCpCmd.Flags().Bool("inline", false, "")
	livewireCpCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(livewireCpCmd)
}
