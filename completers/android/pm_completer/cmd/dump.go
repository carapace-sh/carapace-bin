package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Print various system state associated with the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()

	rootCmd.AddCommand(dumpCmd)

	carapace.Gen(dumpCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
