package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_startCmd = &cobra.Command{
	Use:   "start",
	Short: "[experimental] Start project daemons and the daemons they depend on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_startCmd).Standalone()

	daemons_startCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_startCmd)
}
