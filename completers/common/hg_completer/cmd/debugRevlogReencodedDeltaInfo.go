package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var debugRevlogReencodedDeltaInfoCmd = &cobra.Command{
	Use:    "debug::revlog-reencoded-delta-info",
	Short:  "-c|-m|FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugRevlogReencodedDeltaInfoCmd).Standalone()

	debugRevlogReencodedDeltaInfoCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugRevlogReencodedDeltaInfoCmd.Flags().Bool("delete", false, "delete the result")
	debugRevlogReencodedDeltaInfoCmd.Flags().String("dir", "", "open directory manifest")
	debugRevlogReencodedDeltaInfoCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	debugRevlogReencodedDeltaInfoCmd.Flags().Bool("report", false, "report statistic about the result")
	debugRevlogReencodedDeltaInfoCmd.Flags().Bool("reuse-stored-delta", false, "reuse stored delta when using the same base")
	debugRevlogReencodedDeltaInfoCmd.Flags().String("start-rev", "", "start at rev")
	debugRevlogReencodedDeltaInfoCmd.Flags().String("stop-rev", "", "stop at rev")
	rootCmd.AddCommand(debugRevlogReencodedDeltaInfoCmd)

	carapace.Gen(debugRevlogReencodedDeltaInfoCmd).FlagCompletion(carapace.ActionMap{
		"delete": hg.ActionShelves(),
	})
}
