package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pre-commit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var migrateConfigCmd = &cobra.Command{
	Use:   "migrate-config",
	Short: "Migrate list configuration to new map configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrateConfigCmd).Standalone()

	migrateConfigCmd.Flags().String("color", "", "Whether to use color in output.")
	migrateConfigCmd.Flags().StringP("config", "c", "", "Path to alternate config file")

	carapace.Gen(migrateConfigCmd).FlagCompletion(carapace.ActionMap{
		"color":  action.ActionColors(),
		"config": carapace.ActionFiles(".yaml", ".yml"),
	})

	rootCmd.AddCommand(migrateConfigCmd)
}
