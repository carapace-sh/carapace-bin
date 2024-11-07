package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:     "unset [OPTIONS] <--user|--repo> <NAME>",
	Short:   "Update config file to unset the given option",
	Aliases: []string{},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_unsetCmd).Standalone()

	config_unsetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_unsetCmd.Flags().Bool("repo", false, "Target the repo-level config")
	config_unsetCmd.Flags().Bool("user", false, "Target the user-level config")
	configCmd.AddCommand(config_unsetCmd)

	carapace.Gen(config_unsetCmd).PositionalCompletion(
		jj.ActionConfigs(true).MultiParts("."),
	)
}
