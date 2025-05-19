package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireStubsCmd = &cobra.Command{
	Use:   "livewire:stubs",
	Short: "Publish Livewire stubs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireStubsCmd).Standalone()

	rootCmd.AddCommand(livewireStubsCmd)
}
