package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "inshellisense",
	Short: "IDE style command line auto complete",
	Long:  "https://github.com/microsoft/inshellisense",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("command", "c", "", "command to use as initial input")
	rootCmd.Flags().StringP("duration", "d", "", "duration of the autocomplete session")
	rootCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.Flags().Bool("history", false, "get the last command execute")
	rootCmd.Flags().StringP("shell", "s", "", "shell to use for command execution")
	rootCmd.Flags().BoolP("version", "v", false, "output the current version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"command":  bridge.ActionCarapaceBin().Split(),
		"duration": carapace.ActionValues("single", "session"),
		"shell":    carapace.ActionValues("bash", "pwsh", "zsh", "fish"),
	})
}
