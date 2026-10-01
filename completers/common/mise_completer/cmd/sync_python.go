package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sync_pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Sync python installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sync_pythonCmd).Standalone()

	syncCmd.AddCommand(sync_pythonCmd)
}
