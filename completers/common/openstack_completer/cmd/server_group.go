package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_groupCmd).Standalone()

	serverCmd.AddCommand(server_groupCmd)
}
