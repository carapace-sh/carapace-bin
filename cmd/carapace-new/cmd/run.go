package cmd

import (
	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return spec.ActionSpec(c.Args[0]).Shift(1)
		}),
	)
}
