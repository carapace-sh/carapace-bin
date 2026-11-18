package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_sessionCmd = &cobra.Command{
	Use:    "session",
	Short:  "`atuin threads`, as it was called before",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_sessionCmd).Standalone()

	ai_sessionCmd.Flags().BoolP("help", "h", false, "Print help")
	aiCmd.AddCommand(ai_sessionCmd)
}
