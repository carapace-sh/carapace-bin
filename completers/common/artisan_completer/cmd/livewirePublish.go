package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livewirePublishCmd = &cobra.Command{
	Use:   "livewire:publish [--assets] [--config] [--pagination]",
	Short: "Publish Livewire configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livewirePublishCmd).Standalone()

	livewirePublishCmd.Flags().Bool("assets", false, "Indicates if Livewire's front-end assets should be published")
	livewirePublishCmd.Flags().Bool("config", false, "Indicates if Livewire's config file should be published")
	livewirePublishCmd.Flags().Bool("pagination", false, "Indicates if Livewire's pagination views should be published")
	rootCmd.AddCommand(livewirePublishCmd)
}
