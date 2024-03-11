package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var unsquashCmd = &cobra.Command{
	Use:     "unsquash",
	Short:   "Move changes from a revision's parent into the revision",
	Aliases: []string{"unamend"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsquashCmd).Standalone()

	unsquashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	unsquashCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to unsquash")
	unsquashCmd.Flags().StringP("revision", "r", "", "")
	rootCmd.AddCommand(unsquashCmd)

	carapace.Gen(unsquashCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})
}
