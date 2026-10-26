package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Sets config options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("absolute", false, "Returns absolute paths when fetching *-dir config values instead of relative")
	configCmd.Flags().Bool("append", false, "When adding a repository, append it (lowest priority) to the existing ones instead of prepending it (highest priority)")
	configCmd.Flags().BoolP("auth", "a", false, "Affect auth config file (only used for --editor)")
	configCmd.Flags().BoolP("editor", "e", false, "Open editor")
	configCmd.Flags().StringP("file", "f", "", "If you want to choose a different composer.json or config.json")
	configCmd.Flags().BoolP("global", "g", false, "Apply command to the global config file")
	configCmd.Flags().BoolP("json", "j", false, "JSON decode the setting value, to be used with extra.* keys")
	configCmd.Flags().BoolP("list", "l", false, "List configuration settings")
	configCmd.Flags().BoolP("merge", "m", false, "Merge the setting value with the current value, to be used with extra.* or audit.ignore[-abandoned] keys in combination with --json")
	configCmd.Flags().Bool("source", false, "Display where the config value is loaded from")
	configCmd.Flags().Bool("unset", false, "Unset the given setting-key")

	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(".json"),
	})

	carapace.Gen(configCmd).PositionalCompletion(
		action.ActionConfigKeys(configCmd),
	)

	carapace.Gen(configCmd).PositionalAnyCompletion(
		action.ActionConfigValues(),
	)
}
