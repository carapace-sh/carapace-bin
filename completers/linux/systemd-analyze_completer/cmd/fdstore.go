package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fdstoreCmd = &cobra.Command{
	Use:   "fdstore",
	Short: "Show file descriptor store contents of service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fdstoreCmd).Standalone()

	rootCmd.AddCommand(fdstoreCmd)
	carapace.Gen(fdstoreCmd).PositionalAnyCompletion(
		action.ActionServices(fdstoreCmd).FilterArgs(),
	)
}
