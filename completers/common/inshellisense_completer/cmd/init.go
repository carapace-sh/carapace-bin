package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [options] [shell]",
	Short: "generates shell configurations and prints the source command for the specified shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("generate-full-configs", false, "")
	initCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues("bash", "pwsh", "zsh", "fish", "xonsh", "nu"),
	)
}
