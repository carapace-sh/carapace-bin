package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show server status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverStatusCmd).Standalone()
	serverStatusCmd.Flags().Bool("json", false, "output server information as JSON")

	serverCmd.AddCommand(serverStatusCmd)
}
