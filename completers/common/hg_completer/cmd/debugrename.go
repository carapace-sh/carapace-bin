package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugrenameCmd = &cobra.Command{
	Use:    "debugrename",
	Short:  "dump rename information",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrenameCmd).Standalone()

	debugrenameCmd.Flags().StringP("rev", "r", "", "revision to debug")
	rootCmd.AddCommand(debugrenameCmd)

	carapace.Gen(debugrenameCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
