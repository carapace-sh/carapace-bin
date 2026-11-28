package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var embedModels_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Show or set the default embedding model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(embedModels_defaultCmd).Standalone()

	embedModels_defaultCmd.Flags().Bool("remove-default", false, "Reset to specifying no default model")
	embedModelsCmd.AddCommand(embedModels_defaultCmd)
}
