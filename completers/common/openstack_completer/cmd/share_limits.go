package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_limitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_limitsCmd).Standalone()

	shareCmd.AddCommand(share_limitsCmd)
}
