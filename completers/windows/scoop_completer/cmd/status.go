package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show status and check for new app versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()
	statusCmd.Flags().BoolP("local", "l", false, "check only locally installed apps, disable remote fetching")
	rootCmd.AddCommand(statusCmd)
}
