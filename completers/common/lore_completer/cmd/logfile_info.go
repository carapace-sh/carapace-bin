package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logfile_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logfile_infoCmd).Standalone()

	logfile_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	logfileCmd.AddCommand(logfile_infoCmd)
}
