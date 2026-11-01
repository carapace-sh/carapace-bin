package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "[experimental] Show the output of project daemons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_logsCmd).Standalone()

	daemons_logsCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_logsCmd)
}
