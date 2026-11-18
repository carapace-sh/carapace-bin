package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_resumeCmd = &cobra.Command{
	Use:    "resume",
	Short:  "`atuin threads resume`, as it was called before",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_resumeCmd).Standalone()

	help_aiCmd.AddCommand(help_ai_resumeCmd)
}
