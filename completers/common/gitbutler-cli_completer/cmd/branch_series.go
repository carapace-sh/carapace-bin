package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_seriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Create a new series on top of the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_seriesCmd).Standalone()

	branch_seriesCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_seriesCmd.Flags().StringP("series-name", "s", "", "The name of the series to create on top of the stack")
	branch_seriesCmd.MarkFlagRequired("series-name")
	branchCmd.AddCommand(branch_seriesCmd)
}
