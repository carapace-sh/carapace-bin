package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var hook_runCmd = &cobra.Command{
	Use:   "run [--ignore-missing] [--to-stdin=<path>] <hook-name> [-- <hook-args>]",
	Short: "Run git hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_runCmd).Standalone()

	hook_runCmd.Flags().Bool("ignore-missing", false, "ignore any missing hook by quietly returning zero")
	hook_runCmd.Flags().String("to-stdin", "", "specify a file which will be streamed into the hook’s stdin")
	hookCmd.AddCommand(hook_runCmd)

	carapace.Gen(hook_runCmd).FlagCompletion(carapace.ActionMap{
		"to-stdin": carapace.ActionFiles(),
	})

	carapace.Gen(hook_runCmd).PositionalCompletion(
		git.ActionHooks(),
	)
}
