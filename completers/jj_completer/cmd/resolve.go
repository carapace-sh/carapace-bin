package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve a conflicted file with an external merge tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resolveCmd).Standalone()

	resolveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resolveCmd.Flags().BoolP("list", "l", false, "Instead of resolving one conflict, list all the conflicts")
	resolveCmd.Flags().BoolP("quiet", "q", false, "Do not print the list of remaining conflicts (if any) after resolving a conflict")
	resolveCmd.Flags().StringP("revision", "r", "@", "")
	rootCmd.AddCommand(resolveCmd)

	carapace.Gen(resolveCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(resolveCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionConflicts(resolveCmd.Flag("revision").Value.String()).FilterArgs()
		}),
	)
}
