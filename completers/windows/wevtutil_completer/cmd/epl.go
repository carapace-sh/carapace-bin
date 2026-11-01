package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eplCmd = &cobra.Command{
	Use:   "epl",
	Short: "export events to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eplCmd).Standalone()
	rootCmd.AddCommand(eplCmd)
}
