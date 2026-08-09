package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/paru_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/paru"
	"github.com/spf13/cobra"
)

var repoctlCmd = &cobra.Command{
	Use:     "repoctl",
	Aliases: []string{"L"},
	Short:   "List local repos",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoctlCmd).Standalone()

	repoctlCmd.Flags().CountP("clean", "c", "Remove packages that are not currently installed from repos")
	repoctlCmd.Flags().CountP("delete", "d", "Remove a package from the local repo (-dd to also uninstall)")
	repoctlCmd.Flags().BoolP("list", "l", false, "List packages in local repos")
	repoctlCmd.Flags().BoolP("quiet", "q", false, "Show less information")
	repoctlCmd.Flags().CountP("refresh", "y", "Refresh local repos")
	common.AddNewFlags(repoctlCmd)

	carapace.Gen(repoctlCmd).PositionalAnyCompletion(
		paru.ActionPackageSearch().FilterArgs(),
	)
}
