package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugDeltaFindCmd = &cobra.Command{
	Use:    "debug-delta-find",
	Short:  "-c|-m|FILE REV",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugDeltaFindCmd).Standalone()

	debugDeltaFindCmd.Flags().BoolP("changelog", "c", false, "open changelog")
	debugDeltaFindCmd.Flags().String("dir", "", "open directory manifest")
	debugDeltaFindCmd.Flags().BoolP("manifest", "m", false, "open manifest")
	debugDeltaFindCmd.Flags().String("source", "", "input data feed to the process (full, storage, p1, p2, prev)")
	debugDeltaFindCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugDeltaFindCmd)
}
