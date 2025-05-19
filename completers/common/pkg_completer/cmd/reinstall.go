package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Reinstall specified installed packages at the latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)

	rootCmd.AddCommand(reinstallCmd)
}
