package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Read, set, or list configuration values",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().StringP("file", "f", "", "Read or write only the given config file (no repository required)")
	configCmd.Flags().Bool("global", false, "Use the global (per-user) config file instead of this repository's")
	configCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.Flags().BoolP("list", "l", false, "List all configuration values")
	configCmd.Flags().Bool("local", false, "Read or write only this repository's config file (`.git/config`)")
	configCmd.Flags().Bool("unset", false, "Remove the key instead of reading or setting it")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(configCmd).PositionalCompletion(
		action.ActionConfigKeys(),
		carapace.ActionValues(),
	)
}
