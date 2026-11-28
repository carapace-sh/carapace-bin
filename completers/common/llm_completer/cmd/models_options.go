package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "Manage default options for models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_optionsCmd).Standalone()

	modelsCmd.AddCommand(models_optionsCmd)
}
