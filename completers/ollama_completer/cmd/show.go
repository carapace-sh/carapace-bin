package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ollama"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show MODEL",
	Short: "Show information for a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().Bool("license", false, "Show license of a model")
	showCmd.Flags().Bool("modelfile", false, "Show Modelfile of a model")
	showCmd.Flags().Bool("parameters", false, "Show parameters of a model")
	showCmd.Flags().Bool("system", false, "Show system message of a model")
	showCmd.Flags().Bool("template", false, "Show template of a model")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalCompletion(
		ollama.ActionModels(),
	)
}
