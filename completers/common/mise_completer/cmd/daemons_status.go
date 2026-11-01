package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "[experimental] Show the status of project daemons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_statusCmd).Standalone()

	daemons_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_statusCmd)
}
