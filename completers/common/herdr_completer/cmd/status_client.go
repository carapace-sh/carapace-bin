package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var status_clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Show local client status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(status_clientCmd).Standalone()

	status_clientCmd.Flags().Bool("json", false, "")
	statusCmd.AddCommand(status_clientCmd)
}
