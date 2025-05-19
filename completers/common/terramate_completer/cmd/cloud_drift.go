package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud_driftCmd = &cobra.Command{
	Use:   "drift",
	Short: "manage cloud drifts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_driftCmd).Standalone()

	cloudCmd.AddCommand(cloud_driftCmd)
}
