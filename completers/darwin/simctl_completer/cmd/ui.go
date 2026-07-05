package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Get or set UI options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uiCmd).Standalone()
	rootCmd.AddCommand(uiCmd)
	carapace.Gen(uiCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("appearance"),
		carapace.ActionValues("light", "dark"),
	)
}
