package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var unsignCmd = &cobra.Command{
	Use:   "unsign [OPTIONS]",
	Short: "Drop a cryptographic signature",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsignCmd).Standalone()

	unsignCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	unsignCmd.Flags().StringSliceP("revisions", "r", []string{"@"}, "What revision(s) to unsign")
	rootCmd.AddCommand(unsignCmd)

	carapace.Gen(unsignCmd).FlagCompletion(carapace.ActionMap{
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
