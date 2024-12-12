package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "set-env <name> <value>",
	Short:              "set environment variable",
	Long:               "https://carapace-sh.github.io/carapace-bin/environment.html",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		env.ActionNameValues(true),
	)
}
