package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var architecturesCmd = &cobra.Command{
	Use:   "architectures",
	Short: "List known architectures",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(architecturesCmd).Standalone()

	rootCmd.AddCommand(architecturesCmd)

	carapace.Gen(architecturesCmd).PositionalAnyCompletion(
		action.ActionArchitectures(),
	)
}
