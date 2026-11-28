package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reconfigureCmd = &cobra.Command{
	Use:   "reconfigure",
	Short: "Re-run CMake",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconfigureCmd).Standalone()

	rootCmd.AddCommand(reconfigureCmd)
}
