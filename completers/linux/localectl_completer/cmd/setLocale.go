package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/localectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setLocaleCmd = &cobra.Command{
	Use:   "set-locale",
	Short: "Set system locale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setLocaleCmd).Standalone()

	rootCmd.AddCommand(setLocaleCmd)

	carapace.Gen(setLocaleCmd).PositionalAnyCompletion(
		action.ActionLocales().FilterArgs(),
	)
}
