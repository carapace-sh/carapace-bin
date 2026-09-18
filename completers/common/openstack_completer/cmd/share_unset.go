package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset share properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_unsetCmd).Standalone()

	share_unsetCmd.Flags().Bool("description", false, "Unset share description.")
	share_unsetCmd.Flags().Bool("name", false, "Unset share name.")
	share_unsetCmd.Flags().String("property", "", "Remove a property from share (repeat option to remove multiple properties)")
	shareCmd.AddCommand(share_unsetCmd)
}
