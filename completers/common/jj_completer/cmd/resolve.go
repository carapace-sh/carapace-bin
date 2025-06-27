package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve conflicted files with an external merge tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.Flags().BoolP("list", "l", false, "Instead of resolving conflicts, list all the conflicts")
	resolveCmd.Flags().StringP("revision", "r", "@", "")
	resolveCmd.Flags().String("tool", "", "Specify 3-way merge tool to be used")
	rootCmd.AddCommand(resolveCmd)

	carapace.Gen(resolveCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"tool":     bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(resolveCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionConflicts(resolveCmd.Flag("revision").Value.String()).FilterArgs()
		}),
	)
}
