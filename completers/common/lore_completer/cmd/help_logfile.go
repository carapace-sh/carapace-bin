package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_logfileCmd = &cobra.Command{
	Use:   "logfile",
	Short: "Logfile commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_logfileCmd).Standalone()

	helpCmd.AddCommand(help_logfileCmd)
}
