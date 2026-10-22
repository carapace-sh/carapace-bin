package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_accessCmd = &cobra.Command{
	Use:   "access",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_accessCmd).Standalone()

	share_typeCmd.AddCommand(share_type_accessCmd)
}
