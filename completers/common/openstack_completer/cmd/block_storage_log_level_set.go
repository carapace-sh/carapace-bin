package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var block_storage_log_level_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set log level of block storage service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(block_storage_log_level_setCmd).Standalone()

	block_storage_log_level_setCmd.Flags().String("host", "", "Set block storage service log level of specified host (name only)")
	block_storage_log_level_setCmd.Flags().String("log-prefix", "", "Prefix for the log, e.g. 'sqlalchemy'")
	block_storage_log_level_setCmd.Flags().String("service", "", "Set block storage service log level of specified service (name only)")
	block_storage_log_levelCmd.AddCommand(block_storage_log_level_setCmd)
}
