package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var unfreezeCmd = &cobra.Command{
	Use:   "unfreeze",
	Short: "Unfreeze a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unfreezeCmd).Standalone()

	unfreezeCmd.Flags().Bool("sticky", false, "persists the unfrozen state for the process lifetime")

	rootCmd.AddCommand(unfreezeCmd)

	carapace.Gen(unfreezeCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
