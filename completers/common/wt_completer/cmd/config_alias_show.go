package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_alias_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show an alias's template, or all aliases' templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_alias_showCmd).Standalone()

	config_alias_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_aliasCmd.AddCommand(config_alias_showCmd)

	carapace.Gen(config_alias_showCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
