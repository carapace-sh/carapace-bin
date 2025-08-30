package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "inshellisense [options] [command]",
	Short: "IDE style command line auto complete",
	Long:  "https://github.com/microsoft/inshellisense",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("check", "c", false, "check if shell is in an inshellisense session")
	rootCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.Flags().BoolP("login", "l", false, "start shell as a login shell")
	rootCmd.Flags().StringP("shell", "s", "", "shell to use for command execution, supported shells:")
	rootCmd.Flags().BoolP("verbose", "V", false, "enable verbose logging")
	rootCmd.Flags().BoolP("version", "v", false, "output the current version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "pwsh", "zsh", "fish", "xonsh", "nu"),
	})
}
