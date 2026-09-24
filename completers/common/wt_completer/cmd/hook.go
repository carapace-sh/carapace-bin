package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Run configured hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookCmd).Standalone()

	hookCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(hookCmd)

	carapace.Gen(hookCmd).PositionalCompletion(
		wt.ActionHookTypes(),
	)

	carapace.Gen(hookCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			return wt.ActionHookCommands(c.Args[0])
		}),
	)
}
