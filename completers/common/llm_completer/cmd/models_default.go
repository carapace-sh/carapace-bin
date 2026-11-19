package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Show or set the default model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_defaultCmd).Standalone()

	modelsCmd.AddCommand(models_defaultCmd)
}
