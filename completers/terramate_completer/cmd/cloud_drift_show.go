package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud_drift_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show drifts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_drift_showCmd).Standalone()

	cloud_driftCmd.AddCommand(cloud_drift_showCmd)
}
