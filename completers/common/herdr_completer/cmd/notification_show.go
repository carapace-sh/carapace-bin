package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notification_showCmd = &cobra.Command{
	Use:   "show <title>",
	Short: "Show a notification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notification_showCmd).Standalone()

	notification_showCmd.Flags().String("body", "", "")
	notification_showCmd.Flags().String("position", "", "")
	notification_showCmd.Flags().String("sound", "", "")
	notificationCmd.AddCommand(notification_showCmd)

	carapace.Gen(notification_showCmd).FlagCompletion(carapace.ActionMap{
		"position": carapace.ActionValues("top-left", "top-right", "bottom-left", "bottom-right"),
		"sound":    carapace.ActionValues("none", "done", "request"),
	})
}
