package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "set-env <name> <value>",
	Short: "Set an environment variable",
	Long:  "https://elv.sh/ref/builtin.html#set-env",
	DisableFlagParsing: true,
	Run:   func(cmd *cobra.Command, args []string) {},
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
