package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewireFormCmd = &cobra.Command{
	Use:   "livewire:form [--force] [--] <name>",
	Short: "Create a new Livewire form class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewireFormCmd).Standalone()

	livewireFormCmd.Flags().Bool("force", false, "")
	rootCmd.AddCommand(livewireFormCmd)
}
