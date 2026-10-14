package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_propertiesCmd = &cobra.Command{
	Use:   "properties",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_propertiesCmd).Standalone()

	shareCmd.AddCommand(share_propertiesCmd)
}
