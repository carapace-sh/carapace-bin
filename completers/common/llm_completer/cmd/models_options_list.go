package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var models_options_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List default options for all models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(models_options_listCmd).Standalone()

	models_optionsCmd.AddCommand(models_options_listCmd)
}
