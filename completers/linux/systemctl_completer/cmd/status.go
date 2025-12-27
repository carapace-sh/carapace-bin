package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show runtime status of one or more units",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			opts := systemctl.UnitFileOpts{}.Default()
			opts.User = statusCmd.Root().Flag("user").Changed
			return systemctl.ActionUnitFiles(opts).FilterArgs()
		}),
	)
}
