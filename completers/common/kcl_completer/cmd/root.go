package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "kcl",
	Short:              "KCL command line interface",
	Long:               "https://kcl-lang.io",
	DisableFlagParsing: true,
	Run:                func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		bridge.ActionCobra("kcl"),
	)
}
