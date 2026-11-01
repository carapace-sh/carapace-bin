package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trunkCmd = &cobra.Command{
	Use:     "trunk",
	Short:   "Check out the trunk branch of the stack",
	GroupID: "nav",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trunkCmd).Standalone()

	rootCmd.AddCommand(trunkCmd)
}
