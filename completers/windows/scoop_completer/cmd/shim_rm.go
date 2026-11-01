package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
	"github.com/spf13/cobra"
)

var shim_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shim_rmCmd).Standalone()
	shimCmd.AddCommand(shim_rmCmd)

	carapace.Gen(shim_rmCmd).PositionalAnyCompletion(
		scoop.ActionShims(),
	)
}
