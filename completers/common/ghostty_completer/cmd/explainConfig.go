package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ghostty"
	"github.com/spf13/cobra"
)

var explainConfigCmd = &cobra.Command{
	Use:   "+explain-config",
	Short: "Explain a config key or action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainConfigCmd).Standalone()

	explainConfigCmd.Flags().Bool("help", false, "show help")
	explainConfigCmd.Flags().String("keybind", "", "The name of the keybind action to explain")
	explainConfigCmd.Flags().Bool("no-pager", false, "Disable automatic paging of output")
	explainConfigCmd.Flags().String("option", "", "The name of the configuration option to explain")
	rootCmd.AddCommand(explainConfigCmd)

	carapace.Gen(explainConfigCmd).FlagCompletion(carapace.ActionMap{
		"keybind": ghostty.ActionKeybindActions(),
		"option":  ghostty.ActionConfigKeys(),
	})
}
