package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline_logCmd = &cobra.Command{
	Use:   "log",
	Short: "see logs from offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline_logCmd).Standalone()

	offline_logCmd.Flags().Int("number", 0, "show the log specified by the boot number")
	offlineCmd.AddCommand(offline_logCmd)
}
