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
	rootCmd.Flags().StringP("flags", "s", "", "Set the flags for the script")
	rootCmd.Flags().BoolP("interactive", "I", false, "Interactive mode")
	rootCmd.Flags().StringP("language", "l", "", "Override the language for any plain text scripts")
	rootCmd.Flags().StringP("outputmode", "o", "", "Set the output mode")
	rootCmd.Flags().BoolP("stdin", "i", false, "Read from stdin if no script is provided")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"language":   carapace.ActionValues("AppleScript", "JavaScript", "JXA", "Generic"),
		"outputmode": carapace.ActionValues("text", "list", "record"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
