package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var debugp1copiesCmd = &cobra.Command{
	Use:    "debugp1copies",
	Short:  "dump copy information compared to p1",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugp1copiesCmd).Standalone()

	debugp1copiesCmd.Flags().StringP("rev", "r", "", "revision to debug")
	rootCmd.AddCommand(debugp1copiesCmd)

	carapace.Gen(debugp1copiesCmd).FlagCompletion(carapace.ActionMap{
		"rev": hg.ActionRevisions(),
	})
}
