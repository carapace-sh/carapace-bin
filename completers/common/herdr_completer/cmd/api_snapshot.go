package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var api_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Print the live session snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(api_snapshotCmd).Standalone()

	apiCmd.AddCommand(api_snapshotCmd)
}
