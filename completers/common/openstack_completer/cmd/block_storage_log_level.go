package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_log_levelCmd = &cobra.Command{
	Use:   "level",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_log_levelCmd).Standalone()

	block_storage_logCmd.AddCommand(block_storage_log_levelCmd)
}
