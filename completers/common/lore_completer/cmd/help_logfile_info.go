package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_logfile_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_logfile_infoCmd).Standalone()

	help_logfileCmd.AddCommand(help_logfile_infoCmd)
}
