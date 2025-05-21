package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print and set up required environment variables for fnm",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(envCmd)

	addCommonFlags(envCmd)

	envCmd.Flags().Bool("json", false, "Print JSON instead of shell commands")
	envCmd.Flags().String("shell", "", "The shell syntax to use. Infers when missing")
	envCmd.Flags().Bool("use-on-cd", false, "Print the script to change Node versions every directory change")

	carapace.Gen(envCmd).Standalone()

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "zsh", "fish", "powershell"),
	})
}
