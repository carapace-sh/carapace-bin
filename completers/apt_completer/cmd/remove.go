package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apt_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [pattern]...",
	Short: "remove package binaries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	common.AddGetFlags(removeCmd)
	common.ActionInstallFlags(removeCmd)
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		apt.ActionPackages(),
	)
}
