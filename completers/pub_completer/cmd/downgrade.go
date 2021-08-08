package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pub_completer/cmd/action"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade the current package's dependencies to oldest versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	downgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	downgradeCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	downgradeCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	rootCmd.AddCommand(downgradeCmd)

	carapace.Gen(downgradeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDependencies().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
