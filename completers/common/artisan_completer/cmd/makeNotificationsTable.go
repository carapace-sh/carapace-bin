package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeNotificationsTableCmd = &cobra.Command{
	Use:     "make:notifications-table",
	Short:   "Create a migration for the notifications table",
	Aliases: []string{"notifications:table"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeNotificationsTableCmd).Standalone()

	rootCmd.AddCommand(makeNotificationsTableCmd)
}
