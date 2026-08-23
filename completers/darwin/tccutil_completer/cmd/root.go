package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tccutil",
	Short: "manage the privacy database",
	Long:  "https://keith.github.io/xcode-manpages/tccutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("reset"),
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"AddressBook", "Address book",
			"All", "All services",
			"Calendar", "Calendar",
			"Camera", "Camera",
			"ContactsFull", "Contacts full access",
			"ContactsLimited", "Contacts limited access",
			"DeveloperTool", "Developer tool",
			"Facebook", "Facebook",
			"LinkedIn", "LinkedIn",
			"Liverpool", "Liverpool",
			"Location", "Location services",
			"Microphone", "Microphone",
			"Photos", "Photos",
			"PhotosAdd", "Photos add-only access",
			"Proximity", "Proximity",
			"Reminders", "Reminders",
			"ScreenCapture", "Screen capture",
			"ShareKit", "ShareKit",
			"SinaWeibo", "Sina Weibo",
			"Siri", "Siri",
			"SpeechRecognition", "Speech recognition",
			"TencentWeibo", "Tencent Weibo",
			"Twitter", "Twitter",
			"Ubiquity", "Ubiquity",
			"Undeletable", "Undeletable",
			"Willow", "Willow",
		),
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}