package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var securityCmd = &cobra.Command{
	Use:   "security",
	Short: "Analyze security of unit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityCmd).Standalone()

	rootCmd.AddCommand(securityCmd)

	carapace.Gen(securityCmd).PositionalAnyCompletion(
		action.ActionServices(securityCmd).FilterArgs(),
	)
}
