package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Send notifications to the user",
}

func init() {
	rootCmd.AddCommand(notifyCmd)
	carapace.Gen(notifyCmd).Standalone()

	notifyCmd.Flags().StringArrayP("icon", "n", nil, "The name of the icon to use for the notification. An icon with this name will be searched for on the computer running the terminal emulator. Can be specified multiple times, the first name that is found will be used. Standard names: error, file-manager, help, info, question, system-monitor, text-editor, warn, warning")
	notifyCmd.Flags().StringP("icon-path", "p", "", "Path to an image file in PNG/JPEG/GIF formats to use as the icon. If both name and path are specified then first the name will be looked for and if not found then the path will be used.")
	notifyCmd.Flags().StringP("app-name", "a", "", "The application name for the notification.")
	notifyCmd.Flags().StringArrayP("button", "b", nil, "Add a button with the specified text to the notification. Can be specified multiple times for multiple buttons. If --wait-till-closed is used then the kitten will print the button number to STDOUT if the user clicks a button. 1 for the first button, 2 for the second button and so on.")
	notifyCmd.Flags().StringP("urgency", "u", "", "The urgency of the notification.")
	notifyCmd.Flags().StringP("expire-after", "e", "", "The duration, for the notification to appear on screen. The default is to use the policy of the OS notification service. A value of never means the notification should never expire, however, this may or may not work depending on the policies of the OS notification service. Time is specified in the form NUMBER[SUFFIX] where SUFFIX can be s for seconds, m for minutes, h for hours or d for days.")
	notifyCmd.Flags().StringP("sound-name", "s", "", "The name of the sound to play with the notification. system means let the notification system use whatever sound it wants. silent means prevent any sound from being played. Any other value is passed to the desktop's notification system which may or may not honor it.")
	notifyCmd.Flags().StringP("type", "t", "", "The notification type. Can be any string, it is used by users to create filter rules for notifications, so choose something descriptive of the notification's purpose.")
	notifyCmd.Flags().StringP("identifier", "i", "", "The identifier of this notification. If a notification with the same identifier is already displayed, it is replaced/updated.")
	notifyCmd.Flags().StringP("print-identifier", "P", "", "Print the identifier for the notification to STDOUT. Useful when not specifying your own identifier via the --identifier option.")
	notifyCmd.Flags().BoolP("wait-for-completion", "w", false, "Wait until the notification is closed. If the user activates the notification, \"0\" is printed to STDOUT before quitting. If a button on the notification is pressed the number corresponding to the button is printed to STDOUT. Press the Esc or Ctrl+C keys to close the notification manually.")
	notifyCmd.Flags().Bool("wait-till-closed", false, "Wait until the notification is closed. If the user activates the notification, \"0\" is printed to STDOUT before quitting. If a button on the notification is pressed the number corresponding to the button is printed to STDOUT. Press the Esc or Ctrl+C keys to close the notification manually.")
	notifyCmd.Flags().String("wait-till-closed", "", "Wait until the notification is closed. If the user activates the notification, \"0\" is printed to STDOUT before quitting. If a button on the notification is pressed the number corresponding to the button is printed to STDOUT. Press the Esc or Ctrl+C keys to close the notification manually.")
	notifyCmd.Flags().Bool("only-print-escape-code", false, "Only print the escape code to STDOUT. Useful if using this kitten as part of a larger application. If this is specified, the --wait-till-closed option will be used for escape code generation, but no actual waiting will be done.")
	notifyCmd.Flags().StringP("icon-cache-id", "g", "", "Identifier to use when caching icons in the terminal emulator. Using an identifier means that icon data needs to be transmitted only once using --icon-path. Subsequent invocations will use the cached icon data, at least until the terminal instance is restarted. This is useful if this kitten is being used inside a larger application, with --only-print-escape-code.")
	notifyCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(notifyCmd).FlagCompletion(carapace.ActionMap{
		"icon":       carapace.ActionValues("error", "file-manager", "help", "info", "question", "system-monitor", "text-editor", "warn", "warning"),
		"icon-path":  carapace.ActionFiles("png", "jpeg", "jpg", "gif"),
		"urgency":    carapace.ActionValues("normal", "critical", "low"),
		"sound-name": carapace.ActionValues("system", "silent"),
	})
}
