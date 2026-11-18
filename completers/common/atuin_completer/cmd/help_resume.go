package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_resumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "Pick a captured coding-agent thread and resume it (short for `atuin threads resume`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_resumeCmd).Standalone()

	helpCmd.AddCommand(help_resumeCmd)
}
