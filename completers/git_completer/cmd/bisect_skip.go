package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bisect_skipCmd = &cobra.Command{
	Use:   "skip",
	Short: "mark revisions as untestable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_skipCmd).Standalone()

	bisectCmd.AddCommand(bisect_skipCmd)

	carapace.Gen(bisect_skipCmd).DashAnyCompletion(
		carapace.ActionPositional(bisect_skipCmd),
	)
}
