package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiPackIndex_expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Delete the pack-files that are tracked by the MIDX file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiPackIndex_expireCmd).Standalone()

	multiPackIndexCmd.AddCommand(multiPackIndex_expireCmd)
}
