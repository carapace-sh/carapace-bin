package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Updates configuration with the value of a given key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("global", "g", false, "Write to global '~/.config/glab-cli/config.yml' file rather than the repository's '.git/glab-cli/config.yml' file.")
	config_setCmd.Flags().String("host", "", "Set per-host setting.")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"host": action.ActionConfigHosts(),
	})

	carapace.Gen(config_setCmd).PositionalCompletion(
		action.ActionConfigKeys(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionConfigValues(c.Args[0])
		}),
	)
}
