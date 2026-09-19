package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var debugp2copiesCmd = &cobra.Command{
	Use:    "debugp2copies",
	Short:  "dump copy information compared to p2",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugp2copiesCmd).Standalone()

	debugp2copiesCmd.Flags().StringP("rev", "r", "", "revision to debug")
	rootCmd.AddCommand(debugp2copiesCmd)

	carapace.Gen(debugp2copiesCmd).FlagCompletion(carapace.ActionMap{
		"rev": hg.ActionRevisions(),
	})
}
