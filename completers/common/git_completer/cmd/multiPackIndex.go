package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiPackIndexCmd = &cobra.Command{
	Use:     "multi-pack-index",
	Short:   "Write and verify multi-pack-indexes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(multiPackIndexCmd).Standalone()

	multiPackIndexCmd.Flags().Bool("no-progress", false, "turn progress off")
	multiPackIndexCmd.Flags().String("object-dir", "", "use given directory for the location of Git objects")
	multiPackIndexCmd.Flags().Bool("progress", false, "turn progress on")
	rootCmd.AddCommand(multiPackIndexCmd)

	carapace.Gen(multiPackIndexCmd).FlagCompletion(carapace.ActionMap{
		"object-dir": carapace.ActionDirectories(),
	})
}
