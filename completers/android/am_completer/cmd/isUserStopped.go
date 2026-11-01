package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var isUserStoppedCmd = &cobra.Command{
	Use:   "is-user-stopped",
	Short: "Return whether USER_ID has been stopped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isUserStoppedCmd).Standalone()

	rootCmd.AddCommand(isUserStoppedCmd)

	carapace.Gen(isUserStoppedCmd).PositionalCompletion(
		android.ActionUsers(),
	)
}
