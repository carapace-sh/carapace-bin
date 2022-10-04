package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gum_completer/cmd/action"
	"github.com/spf13/cobra"
)

var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Ask a user to confirm an action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(confirmCmd).Standalone()

	confirmCmd.Flags().String("affirmative", "", "The title of the affirmative action")
	confirmCmd.Flags().Bool("default", false, "Default confirmation action")
	confirmCmd.Flags().String("negative", "", "The title of the negative action")
	confirmCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	confirmCmd.Flags().String("selected.foreground", "", "Foreground Color")
	confirmCmd.Flags().String("timeout", "", "Timeout for confirmation")
	confirmCmd.Flags().String("unselected.foreground", "", "Foreground Color")
	rootCmd.AddCommand(confirmCmd)

	carapace.Gen(confirmCmd).FlagCompletion(carapace.ActionMap{
		"prompt.foreground":     action.ActionColors(),
		"selected.foreground":   action.ActionColors(),
		"unselected.foreground": action.ActionColors(),
	})
}
