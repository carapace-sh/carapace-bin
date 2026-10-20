package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var unuseCmd = &cobra.Command{
	Use:     "unuse",
	Short:   "Remove tool requests from config files and prune unused installations",
	Aliases: []string{"rm", "remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unuseCmd).Standalone()

	unuseCmd.Flags().StringP("env", "e", "", "Edit mise.<ENV>.toml in the current directory, or .mise.<ENV>.toml if that exists")
	unuseCmd.Flags().String("file", "", "Edit this config file, or the config file in this directory")
	unuseCmd.Flags().BoolP("global", "g", false, "Use the global config file (`~/.config/mise/config.toml`) instead of the local one")
	unuseCmd.Flags().BoolP("help", "h", false, "Print help")
	unuseCmd.Flags().Bool("no-prune", false, "Do not also prune the installed version")
	unuseCmd.Flags().Bool("override", false, "Write to the global override file, `config.override.toml`")
	unuseCmd.Flags().StringP("path", "p", "", "Edit this config file, or the config file in this directory")
	rootCmd.AddCommand(unuseCmd)

	carapace.Gen(unuseCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})

	carapace.Gen(unuseCmd).PositionalAnyCompletion(
		mise.ActionInstalledToolVersions(),
	)
}
