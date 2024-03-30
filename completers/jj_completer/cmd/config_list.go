package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list [OPTIONS] [NAME]",
	Short:   "List variables set in config file, along with their values",
	Aliases: []string{"l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	config_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_listCmd.Flags().Bool("include-defaults", false, "Whether to explicitly include built-in default values in the list")
	config_listCmd.Flags().StringP("template", "T", "", "Render each variable using the given template")
	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionConfigs(config_listCmd.Flag("include-defaults").Changed).MultiParts(".")
		}),
	)
}
