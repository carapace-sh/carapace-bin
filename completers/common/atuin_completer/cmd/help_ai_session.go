package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_sessionCmd = &cobra.Command{
	Use:    "session",
	Short:  "`atuin threads`, as it was called before",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_sessionCmd).Standalone()

	help_aiCmd.AddCommand(help_ai_sessionCmd)
}
