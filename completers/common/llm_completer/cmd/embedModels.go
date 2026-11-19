package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var embedModelsCmd = &cobra.Command{
	Use:   "embed-models",
	Short: "Manage available embedding models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(embedModelsCmd).Standalone()

	rootCmd.AddCommand(embedModelsCmd)
}
