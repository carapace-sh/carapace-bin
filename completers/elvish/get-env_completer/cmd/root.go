package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "get-env <name>",
	Short:              "Get the value of an environment variable",
	Long:               "https://elv.sh/ref/builtin.html#get-env",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetPrefix('&')

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionEnvironmentVariables(),
	)
}
