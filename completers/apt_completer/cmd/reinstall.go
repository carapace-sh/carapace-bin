package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apt_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall [pattern]...",
	Short: "reinstall packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	common.AddGetFlags(reinstallCmd)
	common.ActionInstallFlags(reinstallCmd)
	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		apt.ActionPackages(),
	)
}
