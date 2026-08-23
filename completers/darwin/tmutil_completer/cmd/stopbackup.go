package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stopbackupCmd = &cobra.Command{
	Use:   "stopbackup",
	Short: "cancel a backup in progress",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopbackupCmd).Standalone()
	rootCmd.AddCommand(stopbackupCmd)
}