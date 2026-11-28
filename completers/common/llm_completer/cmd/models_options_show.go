package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_options_showCmd = &cobra.Command{
	Use:   "show",
	Short: "List default options set for a specific model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_options_showCmd).Standalone()

	models_optionsCmd.AddCommand(models_options_showCmd)
}
