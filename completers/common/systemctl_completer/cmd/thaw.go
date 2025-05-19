package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var thawCmd = &cobra.Command{
	Use:     "thaw",
	Short:   "Resume execution of a frozen unit",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(thawCmd).Standalone()

	rootCmd.AddCommand(thawCmd)

	carapace.Gen(thawCmd).PositionalAnyCompletion(
		action.ActionUnits(thawCmd).FilterArgs(),
	)
}
