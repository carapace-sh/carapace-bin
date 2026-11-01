package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectUpdaterCmd = &cobra.Command{
	Use:    "connect-updater",
	Short:  "Teleport Connect updater commands.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectUpdaterCmd).Standalone()

	rootCmd.AddCommand(connectUpdaterCmd)
}
