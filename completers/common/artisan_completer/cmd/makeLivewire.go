package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeLivewireCmd = &cobra.Command{
	Use:   "make:livewire [--force] [--inline] [--test] [--pest] [--stub [STUB]] [--] <name>",
	Short: "Create a new Livewire component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeLivewireCmd).Standalone()

	makeLivewireCmd.Flags().Bool("force", false, "")
	makeLivewireCmd.Flags().Bool("inline", false, "")
	makeLivewireCmd.Flags().Bool("pest", false, "")
	makeLivewireCmd.Flags().String("stub", "", "")
	makeLivewireCmd.Flags().Bool("test", false, "")
	rootCmd.AddCommand(makeLivewireCmd)
}
