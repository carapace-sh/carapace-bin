package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "Add, remove, or list additional nodes",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nodeCmd).Standalone()

	rootCmd.AddCommand(nodeCmd)
}
