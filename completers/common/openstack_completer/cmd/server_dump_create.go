package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_dump_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a dump file in server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_dump_createCmd).Standalone()

	server_dumpCmd.AddCommand(server_dump_createCmd)
}
