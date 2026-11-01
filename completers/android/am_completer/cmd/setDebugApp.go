package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var setDebugAppCmd = &cobra.Command{
	Use:   "set-debug-app",
	Short: "Set application to debug",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setDebugAppCmd).Standalone()

	setDebugAppCmd.Flags().Bool("persistent", false, "retain this value")
	setDebugAppCmd.Flags().BoolP("wait", "w", false, "wait for debugger when application starts")

	rootCmd.AddCommand(setDebugAppCmd)

	carapace.Gen(setDebugAppCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
