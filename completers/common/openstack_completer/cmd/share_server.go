package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_serverCmd).Standalone()

	shareCmd.AddCommand(share_serverCmd)
}
