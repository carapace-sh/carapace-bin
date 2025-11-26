package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe [OPTIONS] [REVISION]",
	Short: "Update the change description or other metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()

	describeCmd.Flags().Bool("edit", false, "Open an editor")
	describeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	describeCmd.Flags().StringSliceP("message", "m", nil, "The change description to use (don't open editor)")
	describeCmd.Flags().Bool("stdin", false, "Read the change description from stdin")
	rootCmd.AddCommand(describeCmd)

	carapace.Gen(describeCmd).PositionalCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()),
	)
}
