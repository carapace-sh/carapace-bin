package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeNotificationCmd = &cobra.Command{
	Use:   "make:notification [-f|--force] [-m|--markdown [MARKDOWN]] [--test] [--pest] [--phpunit] [--] <name>",
	Short: "Create a new notification class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeNotificationCmd).Standalone()

	makeNotificationCmd.Flags().Bool("force", false, "Create the class even if the notification already exists")
	makeNotificationCmd.Flags().String("markdown", "", "Create a new Markdown template for the notification")
	makeNotificationCmd.Flags().Bool("pest", false, "Generate an accompanying Pest test for the Notification")
	makeNotificationCmd.Flags().Bool("phpunit", false, "Generate an accompanying PHPUnit test for the Notification")
	makeNotificationCmd.Flags().Bool("test", false, "Generate an accompanying Test test for the Notification")
	rootCmd.AddCommand(makeNotificationCmd)
}
