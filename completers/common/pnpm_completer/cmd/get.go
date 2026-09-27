package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the config value for the provided key. Shorthand for `pnpm config get`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.PersistentFlags().BoolP("global", "g", false, "Operate on the global config file")
	getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	getCmd.PersistentFlags().Bool("json", false, "Show all types of values in JSON format (not just objects and arrays)")
	getCmd.PersistentFlags().String("location", "", "Which config to read or write: `project` for the project's config, `global` for the global config")
	rootCmd.AddCommand(getCmd)
}
