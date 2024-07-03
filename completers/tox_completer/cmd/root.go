package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tox"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tox",
	Short: "automation project",
	Long:  "https://tox.wiki/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	add_common_flags(rootCmd)

	// TODO: rest of the args for the default legacy subcommand (maybe share with ./legacy.go?)
	// commands for the legacy subcommand, that's run by default if no subcommand
	rootCmd.Flags().StringS("e", "e", "", "enumerate, comma separated (ALL -> all environments, not set -> use <env_list> from config) (default: <env_list>)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// for the legacy subcommand
		"e":  tox.ActionEnvironments().UniqueList(","),
	})
}
