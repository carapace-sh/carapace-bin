package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "show combined config settings from all hgrc files",
	Aliases: []string{"showconfig", "debugconfig"},
	GroupID: groups[group_help].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("edit", "e", false, "edit user config")
	configCmd.Flags().Bool("exp-all-known", false, "show all known config option (EXPERIMENTAL)")
	configCmd.Flags().BoolP("global", "g", false, "edit global config")
	configCmd.Flags().BoolP("local", "l", false, "edit repository config")
	configCmd.Flags().Bool("non-shared", false, "edit non shared config")
	configCmd.Flags().Bool("shared", false, "edit shared source repository config")
	configCmd.Flags().Bool("source", false, "show source of configuration value")
	configCmd.Flags().StringP("template", "T", "", "display with template")
	configCmd.Flags().BoolP("untrusted", "u", false, "show untrusted configuration options")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).PositionalAnyCompletion(
		hg.ActionConfigKeys(),
	)
}
