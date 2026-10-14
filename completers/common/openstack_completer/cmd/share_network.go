package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_networkCmd).Standalone()

	shareCmd.AddCommand(share_networkCmd)
}
