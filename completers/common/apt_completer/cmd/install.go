package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [pattern]...",
	Short: "install packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	common.AddGetFlags(installCmd)
	common.ActionInstallFlags(installCmd)
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if condition.CompletingPath(c) {
				return carapace.ActionFiles()
			}
			return apt.ActionPackageSearch()
		}),
	)
}
