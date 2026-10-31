package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_listCmd).Standalone()

	models_listCmd.Flags().Bool("async", false, "List async models")
	models_listCmd.Flags().Bool("json", false, "Output as JSON")
	models_listCmd.Flags().StringP("model", "m", "", "Specific model IDs")
	models_listCmd.Flags().Bool("options", false, "Show options for each model, if available")
	models_listCmd.Flags().StringP("query", "q", "", "Search for models matching these strings")
	models_listCmd.Flags().Bool("schemas", false, "List models that support schemas")
	models_listCmd.Flags().Bool("tools", false, "List models that support tools")
	modelsCmd.AddCommand(models_listCmd)
}
