package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "run environments",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	addCommonSubcommandFlags(runCmd)
	addPkgOnlyFlags(runCmd)
	addEnvFilteringFlags(runCmd)
	addEnvSelectFlag(runCmd)
	addCommonRunFlags(runCmd)

	rootCmd.AddCommand(runCmd)
}
