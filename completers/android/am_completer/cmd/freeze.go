package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var freezeCmd = &cobra.Command{
	Use:   "freeze",
	Short: "Freeze a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freezeCmd).Standalone()

	freezeCmd.Flags().Bool("sticky", false, "persists the frozen state for the process lifetime")

	rootCmd.AddCommand(freezeCmd)

	carapace.Gen(freezeCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
