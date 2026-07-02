package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command across a set of revisions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().Bool("clean", false, "Delete each working copy before running the command")
	runCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	runCmd.Flags().StringP("jobs", "j", "", "How many processes should run in parallel")
	runCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	runCmd.Flags().StringSliceP("revision", "r", nil, "The revisions to change")
	runCmd.Flags().StringSlice("revisions", nil, "The revisions to change")
	runCmd.Flags().Bool("root", false, "Run the command from the working-copy root in each commit instead of from the subdirectory `jj run` was invoked from")
	runCmd.Flags().BoolP("x", "x", false, "A no-op option to match the interface of `git rebase -x`")
	runCmd.Flag("revisions").Hidden = true
	runCmd.Flag("x").Hidden = true
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"revision":  jj.ActionRevsets(jj.RevOpts{}.Default()),
		"revisions": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	carapace.Gen(runCmd).DashAnyCompletion(
		carapace.ActionPositional(runCmd),
	)
}
