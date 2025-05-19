package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireLayoutCmd = &cobra.Command{
	Use:   "livewire:layout [--force] [--stub [STUB]]",
	Short: "Create a new app layout file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireLayoutCmd).Standalone()

	livewireLayoutCmd.Flags().Bool("force", false, "")
	livewireLayoutCmd.Flags().String("stub", "", "If you have several stubs, stored in subfolders")
	rootCmd.AddCommand(livewireLayoutCmd)
}
