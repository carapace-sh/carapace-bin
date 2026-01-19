package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_dumpMetadataCmd = &cobra.Command{
	Use:   "dump-metadata",
	Short: "dump the meta in human-readable format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_dumpMetadataCmd).Standalone()

	debug_dumpMetadataCmd.Flags().String("root", "", "path to state directory")
	debugCmd.AddCommand(debug_dumpMetadataCmd)

	carapace.Gen(debug_dumpMetadataCmd).FlagCompletion(carapace.ActionMap{
		"root": carapace.ActionDirectories(),
	})
}
