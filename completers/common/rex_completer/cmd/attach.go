package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach [SESSION]",
	Short: "Attach to a session and show its oldest visible terminal block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).PositionalCompletion(
		action.ActionSessions(attachCmd),
	)
}
