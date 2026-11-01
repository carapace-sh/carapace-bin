package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Manage the pnpm configuration files",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.PersistentFlags().BoolP("global", "g", false, "Operate on the global config file")
	configCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.PersistentFlags().Bool("json", false, "Show all types of values in JSON format (not just objects and arrays)")
	configCmd.PersistentFlags().String("location", "", "Which config to read or write: `project` for the project's config, `global` for the global config")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"location": carapace.ActionValues("global", "project"),
	})
}
