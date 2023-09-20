package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "get-env <name>",
	Short:              "get environment variable",
	Long:               "https://rsteube.github.io/carapace-bin/environment.html",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionEnvironmentVariables(),
	)
}
