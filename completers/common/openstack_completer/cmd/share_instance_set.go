package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_instance_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Explicitly reset share instance status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_instance_setCmd).Standalone()

	share_instance_setCmd.Flags().String("status", "", "Indicate which state to assign the instance.")
	share_instanceCmd.AddCommand(share_instance_setCmd)
}
