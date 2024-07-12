package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var file_chmodCmd = &cobra.Command{
	Use:   "chmod [OPTIONS] <MODE> <PATHS>...",
	Short: "Sets or removes the executable bit for paths in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_chmodCmd).Standalone()

	file_chmodCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_chmodCmd.Flags().StringP("revision", "r", "", "The revision to update")
	fileCmd.AddCommand(file_chmodCmd)

	carapace.Gen(file_chmodCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(file_chmodCmd).PositionalCompletion(
		carapace.ActionStyledValuesDescribed(
			"n", "normal", style.Default,
			"x", "executable", style.Yellow,
		),
	)

	carapace.Gen(file_chmodCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(file_chmodCmd.Flag("revision").Value.String())
		}),
	)
}
