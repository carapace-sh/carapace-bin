package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the sync status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(statusCmd)
}
