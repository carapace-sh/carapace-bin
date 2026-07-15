package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var status_serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Show running server status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(status_serverCmd).Standalone()

	status_serverCmd.Flags().Bool("json", false, "")
	statusCmd.AddCommand(status_serverCmd)
}
