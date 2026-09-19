package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugpvecCmd = &cobra.Command{
	Use:    "debugpvec",
	Short:  "A B",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugpvecCmd).Standalone()

	rootCmd.AddCommand(debugpvecCmd)
}
