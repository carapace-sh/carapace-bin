package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modelShowCmd = &cobra.Command{
	Use:   "model:show [--database [DATABASE]] [--json] [--] <model>",
	Short: "Show information about an Eloquent model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modelShowCmd).Standalone()

	modelShowCmd.Flags().String("database", "", "The database connection to use")
	modelShowCmd.Flags().Bool("json", false, "Output the model as JSON")
	rootCmd.AddCommand(modelShowCmd)
}
