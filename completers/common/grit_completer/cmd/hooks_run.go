package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var hooks_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a hook by name (passes `--` arguments through to the script)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hooks_runCmd).Standalone()

	hooks_runCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hooksCmd.AddCommand(hooks_runCmd)

	carapace.Gen(hooks_runCmd).PositionalCompletion(
		action.ActionHookNames(),
	)
}
