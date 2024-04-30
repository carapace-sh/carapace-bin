package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var parallelizeCmd = &cobra.Command{
	Use:   "parallelize [OPTIONS] [REVISIONS]..",
	Short: "Parallelize revisions by making them siblings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(parallelizeCmd).Standalone()

	parallelizeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(parallelizeCmd)

	carapace.Gen(parallelizeCmd).PositionalAnyCompletion(
		jj.ActionRevSets(jj.RevOption{}.Default()),
	)
}
