package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Output the path to the logs.db file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_pathCmd).Standalone()

	logsCmd.AddCommand(logs_pathCmd)
}
