package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_poolCmd = &cobra.Command{
	Use:   "pool",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_poolCmd).Standalone()

	shareCmd.AddCommand(share_poolCmd)
}
