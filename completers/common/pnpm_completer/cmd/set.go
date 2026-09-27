package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the config key to the value provided. Shorthand for `pnpm config set`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.PersistentFlags().BoolP("global", "g", false, "Operate on the global config file")
	setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	setCmd.PersistentFlags().Bool("json", false, "Show all types of values in JSON format (not just objects and arrays)")
	setCmd.PersistentFlags().String("location", "", "Which config to read or write: `project` for the project's config, `global` for the global config")
	rootCmd.AddCommand(setCmd)
}
