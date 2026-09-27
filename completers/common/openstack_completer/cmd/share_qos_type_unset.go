package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_qos_type_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset qos type specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_qos_type_unsetCmd).Standalone()

	share_qos_type_unsetCmd.Flags().Bool("description", false, "Unset qos type description.")
	share_qos_type_unsetCmd.Flags().String("spec", "", "Remove specified spec from this qos type")
	share_qos_typeCmd.AddCommand(share_qos_type_unsetCmd)
}
