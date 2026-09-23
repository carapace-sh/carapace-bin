package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var switchUserCmd = &cobra.Command{
	Use:   "switch-user",
	Short: "Switch to put USER_ID in the foreground",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchUserCmd).Standalone()

	rootCmd.AddCommand(switchUserCmd)

	carapace.Gen(switchUserCmd).PositionalCompletion(
		android.ActionUsers(),
	)
}
