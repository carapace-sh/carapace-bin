package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var convergeCmd = &cobra.Command{
	Use:   "converge",
	Short: "Converge divergent changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convergeCmd).Standalone()

	convergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	convergeCmd.Flags().BoolP("interactive", "i", false, "No-op flag to pair with --no-interactive")
	convergeCmd.Flags().Bool("no-interactive", false, "Do not prompt the user for help resolving divergence")
	convergeCmd.Flags().StringSliceP("revision", "r", nil, "The search space to look for divergent revisions")
	convergeCmd.Flags().StringSlice("revisions", nil, "The search space to look for divergent revisions")
	convergeCmd.Flag("interactive").Hidden = true
	convergeCmd.Flag("revisions").Hidden = true
	rootCmd.AddCommand(convergeCmd)

	carapace.Gen(convergeCmd).FlagCompletion(carapace.ActionMap{
		"revision":  jj.ActionRevsets(jj.RevOpts{}.Default()),
		"revisions": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})
}