package cmd

import (
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "aws",
	Short:              "Universal Command Line Interface for Amazon Web Services",
	Long:               "https://aws.amazon.com/cli/",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if _, err := exec.LookPath("carapace-aws"); err == nil {
				return bridge.ActionCarapace("carapace-aws")
			}
			return bridge.ActionAws("aws")
		}),
	)
}
