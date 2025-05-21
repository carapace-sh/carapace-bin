package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Set a version as the default version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(defaultCmd).Standalone()

	rootCmd.AddCommand(defaultCmd)

	carapace.Gen(defaultCmd).PositionalAnyCompletion(
		action.ActionLocalVersions(),
	)
}
