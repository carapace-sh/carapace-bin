package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var partitionTableFlashCmd = &cobra.Command{
	Use:   "partition-table-flash",
	Short: "Flash partition table only",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partitionTableFlashCmd).Standalone()

	common.AddFlashFlags(partitionTableFlashCmd)
	rootCmd.AddCommand(partitionTableFlashCmd)
}
