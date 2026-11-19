package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_options_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear default option(s) for a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_options_clearCmd).Standalone()

	models_optionsCmd.AddCommand(models_options_clearCmd)
}
