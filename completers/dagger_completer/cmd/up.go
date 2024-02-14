package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:    "up",
	Short:  "Start a service and expose its ports to the host",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	rootCmd.AddCommand(upCmd)
}
