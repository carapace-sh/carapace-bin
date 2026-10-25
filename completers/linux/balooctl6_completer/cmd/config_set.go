package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the value of a config parameter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).PositionalCompletion(
		carapace.ActionValues("contentIndexing", "hidden"),
		carapace.ActionValues("0", "1", "false", "n", "no", "true", "y", "yes"),
	)
}
