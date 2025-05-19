package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireMakeCmd = &cobra.Command{
	Use:   "livewire:make [--force] [--inline] [--test] [--pest] [--stub [STUB]] [--] <name>",
	Short: "Create a new Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireMakeCmd).Standalone()

	livewireMakeCmd.Flags().Bool("force", false, "")
	livewireMakeCmd.Flags().Bool("inline", false, "")
	livewireMakeCmd.Flags().Bool("pest", false, "")
	livewireMakeCmd.Flags().String("stub", "", "If you have several stubs, stored in subfolders")
	livewireMakeCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(livewireMakeCmd)
}
