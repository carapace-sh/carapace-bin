package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bisect_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a given command to find the first bad revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_runCmd).Standalone()
	bisect_runCmd.Flags().SetInterspersed(false)

	bisect_runCmd.Flags().Bool("find-good", false, "Find the first good revision instead")
	bisect_runCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bisect_runCmd.Flags().StringSliceP("range", "r", nil, "Range of revisions to bisect (can be repeated)")
	bisect_runCmd.MarkFlagRequired("range")
	bisectCmd.AddCommand(bisect_runCmd)

	carapace.Gen(bisect_runCmd).FlagCompletion(carapace.ActionMap{
		"range": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})

	carapace.Gen(bisect_runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
