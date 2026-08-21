package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_volumeCmd).Standalone()

	serverCmd.AddCommand(server_volumeCmd)
}
