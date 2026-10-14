package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_shareCmd = &cobra.Command{
	Use:   "share",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_shareCmd).Standalone()

	serverCmd.AddCommand(server_shareCmd)
}
