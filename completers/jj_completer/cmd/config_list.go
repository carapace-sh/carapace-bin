package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List variables set in config file, along with their values",
	Aliases: []string{"l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	config_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_listCmd.Flags().Bool("include-defaults", false, "Whether to explicitly include built-in default values in the list")
	configCmd.AddCommand(config_listCmd)
}
