package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "osascript",
	Short: "execute AppleScript and other OSA language scripts",
	Long:  "https://keith.github.io/xcode-manpages/osascript.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("evaluate", "e", "", "Enter one line of a script")
	rootCmd.Flags().StringP("flags", "s", "", "Set the flags for the script (e, h, o, or empty)")
	rootCmd.Flags().BoolP("interactive", "i", false, "Interactive mode")
	rootCmd.Flags().StringP("language", "l", "", "Override the language for any plain text scripts")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"flags":    carapace.ActionValues("e", "h", "o", ""),
		"language": carapace.ActionValues("AppleScript", "JavaScript", "JXA", "Generic"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
