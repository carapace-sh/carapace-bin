package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rescueCmd = &cobra.Command{
	Use:     "rescue",
	Short:   "Enter system rescue mode",
	GroupID: "system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rescueCmd).Standalone()

	rootCmd.AddCommand(rescueCmd)
}
