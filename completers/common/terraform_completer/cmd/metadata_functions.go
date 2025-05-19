package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metadata_functionsCmd = &cobra.Command{
	Use:   "functions -json",
	Short: "Show signatures and descriptions for the available functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metadata_functionsCmd).Standalone()

	metadata_functionsCmd.Flags().BoolS("json", "json", false, "Prints out a json representation of the available function signatures")
	metadataCmd.AddCommand(metadata_functionsCmd)
}
