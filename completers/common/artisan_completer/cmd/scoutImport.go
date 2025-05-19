package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scoutImportCmd = &cobra.Command{
	Use:   "scout:import [-c|--chunk [CHUNK]] [--] <model>",
	Short: "Import the given model into the search index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scoutImportCmd).Standalone()

	scoutImportCmd.Flags().String("chunk", "", "The number of records to import at a time (Defaults to configuration value: `scout.chunk.searchable`)")
	rootCmd.AddCommand(scoutImportCmd)
}
