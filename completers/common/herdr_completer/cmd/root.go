package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "herdr",
	Short: "terminal workspace manager for AI coding agents",
	Long:  "https://herdr.dev/docs/cli-reference/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("default-config", false, "Print default configuration and exit")
	rootCmd.Flags().Bool("handoff", false, "Opt into live handoff for update or remote attach")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().Bool("no-session", false, "Run monolithically without server/client session mode")
	rootCmd.Flags().String("remote", "", "Attach through SSH to a remote Herdr server")
	rootCmd.Flags().String("remote-keybindings", "", "Choose local or server keybindings for remote attach")
	rootCmd.Flags().String("session", "", "Use or create a named persistent session")
	rootCmd.Flags().BoolP("version", "V", false, "Print version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"remote-keybindings": carapace.ActionValues("local", "server"),
	})
}
