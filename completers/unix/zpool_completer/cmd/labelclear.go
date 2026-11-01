package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var labelclearCmd = &cobra.Command{
	Use:     "labelclear [-f] device",
	Short:   "remove ZFS label from a device",
	GroupID: "io",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(labelclearCmd).Standalone()

	labelclearCmd.Flags().BoolS("f", "f", false, "force label removal")

	rootCmd.AddCommand(labelclearCmd)

	carapace.Gen(labelclearCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
