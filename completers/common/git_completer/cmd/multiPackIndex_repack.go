package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiPackIndex_repackCmd = &cobra.Command{
	Use:   "repack",
	Short: "Create a new pack-file containing objects in small pack-files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiPackIndex_repackCmd).Standalone()

	multiPackIndexCmd.AddCommand(multiPackIndex_repackCmd)
}
