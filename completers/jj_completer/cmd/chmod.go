package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var chmodCmd = &cobra.Command{
	Use:   "chmod",
	Short: "Sets or removes the executable bit for paths in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chmodCmd).Standalone()

	chmodCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	chmodCmd.Flags().StringP("revision", "r", "", "The revision to update")
	rootCmd.AddCommand(chmodCmd)

	carapace.Gen(chmodCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevsOption{}.Default()),
	})

	carapace.Gen(chmodCmd).PositionalCompletion(
		carapace.ActionStyledValuesDescribed(
			"n", "normal", style.Default,
			"x", "executable", style.Yellow,
		),
	)

	carapace.Gen(chmodCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(chmodCmd.Flag("revision").Value.String())
		}),
	)
}
