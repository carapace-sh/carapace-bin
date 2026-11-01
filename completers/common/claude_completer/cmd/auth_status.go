package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show authentication status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_statusCmd).Standalone()

	auth_statusCmd.Flags().BoolP("help", "h", false, "Display help for command")
	auth_statusCmd.Flags().Bool("json", false, "Output as JSON (default)")
	auth_statusCmd.Flags().Bool("text", false, "Output as human-readable text")
	authCmd.AddCommand(auth_statusCmd)
}
