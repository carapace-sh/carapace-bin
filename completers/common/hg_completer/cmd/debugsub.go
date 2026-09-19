package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debugsubCmd = &cobra.Command{
	Use:    "debugsub",
	Short:  "[-r REV] [REV]",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugsubCmd).Standalone()

	debugsubCmd.Flags().StringP("rev", "r", "", "revision to check")
	rootCmd.AddCommand(debugsubCmd)

	carapace.Gen(debugsubCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
