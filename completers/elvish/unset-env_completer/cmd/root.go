package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "unset-env <name>",
	Short:              "Unset an environment variable",
	Long:               "https://elv.sh/ref/builtin.html#unset-env",
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
