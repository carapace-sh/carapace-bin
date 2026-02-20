package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Start a shell in a pixi environment, run `exit` to leave the shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_shellCmd).Standalone()

	helpCmd.AddCommand(help_shellCmd)
}
