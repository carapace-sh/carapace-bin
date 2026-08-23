package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Return current FileVault status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().Bool("device", false, "Specify the device to check")
	statusCmd.Flags().Bool("extended", false, "Show extended information")
	statusCmd.Flags().Bool("verbose", false, "Enable verbose mode")
}
