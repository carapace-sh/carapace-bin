package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:     "export [-af] pool...",
	Short:   "export a pool",
	GroupID: "io",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().BoolS("a", "a", false, "export all pools")
	exportCmd.Flags().BoolS("f", "f", false, "force export")

	rootCmd.AddCommand(exportCmd)

	carapace.Gen(exportCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
