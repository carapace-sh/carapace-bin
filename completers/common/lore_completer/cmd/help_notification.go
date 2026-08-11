package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_notificationCmd = &cobra.Command{
	Use:   "notification",
	Short: "Notifications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_notificationCmd).Standalone()

	helpCmd.AddCommand(help_notificationCmd)
}
