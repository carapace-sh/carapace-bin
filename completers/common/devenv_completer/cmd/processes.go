package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processesCmd = &cobra.Command{
	Use:   "processes",
	Short: "Start or stop processes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processesCmd).Standalone()

	rootCmd.AddCommand(processesCmd)
}
