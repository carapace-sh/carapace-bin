package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get all process status info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	rootCmd.AddCommand(statusCmd)
}
