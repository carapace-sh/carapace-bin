package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/grit_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reflog_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove one or more reflog entries (`ref@{n}`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_deleteCmd).Standalone()

	reflog_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	reflogCmd.AddCommand(reflog_deleteCmd)

	carapace.Gen(reflog_deleteCmd).PositionalAnyCompletion(
		action.ActionReflogSpecs(),
	)
}
