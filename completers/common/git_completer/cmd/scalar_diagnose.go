package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Collect diagnostic information for a Scalar enlistment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_diagnoseCmd).Standalone()

	scalarCmd.AddCommand(scalar_diagnoseCmd)
}
