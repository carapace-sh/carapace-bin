package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_dumpCmd).Standalone()

	serverCmd.AddCommand(server_dumpCmd)
}
