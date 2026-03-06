package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var hook_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show configured hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_showCmd).Standalone()

	hook_showCmd.Flags().Bool("expanded", false, "Show expanded commands with current variables")
	hook_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hookCmd.AddCommand(hook_showCmd)

	carapace.Gen(hook_showCmd).PositionalCompletion(
		wt.ActionHookTypes(),
	)
}
