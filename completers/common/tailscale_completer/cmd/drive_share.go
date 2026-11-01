package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drive_shareCmd = &cobra.Command{
	Use:   "share",
	Short: "[ALPHA] Create or modify a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drive_shareCmd).Standalone()

	driveCmd.AddCommand(drive_shareCmd)

	carapace.Gen(drive_shareCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
