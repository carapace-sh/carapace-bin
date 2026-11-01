package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promote_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of any current pending promotions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promote_statusCmd).Standalone()

	promote_statusCmd.Flags().String("timeout", "", "Time to wait for promotion completion")
	promote_statusCmd.Flags().Bool("yes", false, "Skip confirmation")

	promoteCmd.AddCommand(promote_statusCmd)
}
