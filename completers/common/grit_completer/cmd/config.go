package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read, set, or list configuration values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("global", false, "Use the global (per-user) config file instead of this repository's")
	configCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.Flags().BoolP("list", "l", false, "List all configuration values")
	configCmd.Flags().Bool("unset", false, "Remove the key instead of reading or setting it")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if configCmd.Flag("list").Changed {
				return carapace.ActionValues()
			}
			switch len(c.Args) {
			case 0:
				return git.ActionConfigs()
			case 1:
				if configCmd.Flag("unset").Changed {
					return carapace.ActionValues()
				}
				return git.ActionConfigValues(c.Args[0])
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
