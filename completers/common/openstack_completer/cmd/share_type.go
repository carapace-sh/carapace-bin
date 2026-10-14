package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_typeCmd = &cobra.Command{
	Use:   "type",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_typeCmd).Standalone()

	shareCmd.AddCommand(share_typeCmd)
}
