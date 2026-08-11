package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logfile_help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logfile_help_infoCmd).Standalone()

	logfile_helpCmd.AddCommand(logfile_help_infoCmd)
}
