package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowFocusCmd = &cobra.Command{
	Use:   "focus [WINDOW|next|prev]",
	Short: "Focus a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowFocusCmd).Standalone()

	windowFocusCmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	windowFocusCmd.Flags().Bool("wrap", false, "wrap next and prev at the ends")

	windowCmd.AddCommand(windowFocusCmd)

	carapace.Gen(windowFocusCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(windowFocusCmd),
	})

	carapace.Gen(windowFocusCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValuesDescribed(
				"next", "focus the next window",
				"prev", "focus the previous window",
			),
			action.ActionWindows(windowFocusCmd),
		).ToA(),
	)
}
