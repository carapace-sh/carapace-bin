package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var engineCmd = &cobra.Command{
	Use:     "engine",
	Short:   "Engine (loadable module) information and manipulation",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(engineCmd).Standalone()

	engineCmd.Flags().BoolS("c", "c", false, "List the capabilities of specified engine")
	engineCmd.Flags().StringS("post", "post", "", "Run command against the ENGINE after loading it")
	engineCmd.Flags().StringS("pre", "pre", "", "Run command against the ENGINE before loading it")
	engineCmd.Flags().BoolS("t", "t", false, "Check that specified engine is available")
	engineCmd.Flags().BoolS("tt", "tt", false, "Display error trace for unavailable engines Commands are like \"SO_PATH:/lib/libdriver.so\"")
	engineCmd.Flags().BoolS("v", "v", false, "List 'control commands' For each specified engine")
	engineCmd.Flags().BoolS("vv", "vv", false, "Also display each command's description")
	engineCmd.Flags().BoolS("vvv", "vvv", false, "Also add the input flags for each command")
	engineCmd.Flags().BoolS("vvvv", "vvvv", false, "Also show internal input flags")
	rootCmd.AddCommand(engineCmd)

	carapace.Gen(engineCmd).PositionalAnyCompletion(
		action.ActionEngines(),
	)
}
