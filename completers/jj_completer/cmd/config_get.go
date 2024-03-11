package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:     "get [OPTIONS] <NAME>",
	Short:   "Get the value of a given config option.",
	Aliases: []string{"g"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		jj.ActionConfigs(true).MultiParts("."),
	)
}
