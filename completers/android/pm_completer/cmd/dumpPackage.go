package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var dumpPackageCmd = &cobra.Command{
	Use:   "dump-package",
	Short: "Print package manager state associated with the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpPackageCmd).Standalone()

	rootCmd.AddCommand(dumpPackageCmd)

	carapace.Gen(dumpPackageCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
