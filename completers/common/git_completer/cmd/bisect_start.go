package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var bisect_startCmd = &cobra.Command{
	Use:   "start",
	Short: "reset bisect state and start bisection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	bisect_startCmd.Flags().Bool("first-parent", false, "follow only the first parent commit upon seeing a merge commit")
	bisect_startCmd.Flags().Bool("no-checkout", false, "do not checkout the new working tree at each iteration")
	bisect_startCmd.Flags().String("reset-when-found", "", "automatically run \"git bisect reset\" when the first bad commit has been determined")
	bisect_startCmd.Flags().String("term-bad", "", "set term for new/bad")
	bisect_startCmd.Flags().String("term-good", "", "set term for old/good")
	bisect_startCmd.Flags().String("term-new", "", "set term for new/bad")
	bisect_startCmd.Flags().String("term-old", "", "set term for old/good")
	carapace.Gen(bisect_startCmd).Standalone()

	bisect_startCmd.Flag("reset-when-found").NoOptDefVal = " "
	bisect_startCmd.MarkFlagsMutuallyExclusive("reset-when-found", "no-checkout")

	bisectCmd.AddCommand(bisect_startCmd)

	carapace.Gen(bisect_startCmd).FlagCompletion(carapace.ActionMap{
		"reset-when-found": carapace.ActionValuesDescribed(
			"original", "reset to the original HEAD",
			"found", "reset to the first bad commit",
		),
	})

	carapace.Gen(bisect_startCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if bisect_startCmd.Flags().ArgsLenAtDash() != -1 {
				return carapace.ActionFiles()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)

	carapace.Gen(bisect_startCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if bisect_startCmd.Flags().ArgsLenAtDash() == 0 {
				return carapace.ActionValues()
			}
			return git.ActionRefFiles(bisect_startCmd.Flags().Args()[0]).FilterArgs()
		}),
	)
}
