package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_options_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a default option for a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_options_setCmd).Standalone()

	models_optionsCmd.AddCommand(models_options_setCmd)
}
