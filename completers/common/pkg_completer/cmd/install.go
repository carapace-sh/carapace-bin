package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install specified packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
