package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var backoutCmd = &cobra.Command{
	Use:   "backout",
	Short: "Apply the reverse of a revision on top of another revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backoutCmd).Standalone()

	backoutCmd.Flags().StringSliceP("destination", "d", nil, "The revision to apply the reverse changes on top of")
	backoutCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	backoutCmd.Flags().StringP("revisions", "r", "", "The revision to apply the reverse of")
	rootCmd.AddCommand(backoutCmd)

	carapace.Gen(backoutCmd).FlagCompletion(carapace.ActionMap{
		"destination": jj.ActionRevs(jj.RevOption{}.Default()),
		"revisions":   jj.ActionRevs(jj.RevOption{}.Default()),
	})
}
