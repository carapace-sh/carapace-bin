package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/linux/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Compare package versions or perform tests on version strings",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "miscellaneous",
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolP("all", "a", false, "Consider packages from all repository tags")
	versionCmd.Flags().BoolP("check", "c", false, "Check versions for validity")
	versionCmd.Flags().BoolP("indexes", "I", false, "Print the version and description for each repository's index")
	versionCmd.Flags().StringP("limit", "l", "", "Limit to packages with output matching given OPERAND")
	versionCmd.Flags().BoolP("test", "t", false, "Compare two version strings")
	common.AddGlobalFlags(versionCmd)
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch {
			case versionCmd.Flag("check").Changed,
				versionCmd.Flag("test").Changed,
				versionCmd.Flag("indexes").Changed:
				return carapace.ActionValues()
			default:
				return action.ActionPackages(versionCmd, true).FilterArgs()
			}
		}),
	)
}
