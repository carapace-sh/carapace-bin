package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_reconfigureCmd = &cobra.Command{
	Use:   "reconfigure",
	Short: "Reconfigure Scalar enlistment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_reconfigureCmd).Standalone()

	scalar_reconfigureCmd.Flags().Bool("all", false, "reconfigure all enlistments")
	scalar_reconfigureCmd.Flags().String("maintenance", "", "maintenance mode")
	scalarCmd.AddCommand(scalar_reconfigureCmd)

	carapace.Gen(scalar_reconfigureCmd).FlagCompletion(carapace.ActionMap{
		"maintenance": carapace.ActionValues("enable", "disable", "keep"),
	})
}
