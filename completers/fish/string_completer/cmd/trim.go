package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trimCmd = &cobra.Command{
	Use:   "trim",
	Short: "Trim leading and trailing characters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trimCmd).Standalone()

	trimCmd.Flags().StringP("chars", "c", "", "set of characters to remove")
	trimCmd.Flags().BoolP("left", "l", false, "only trim leading characters")
	trimCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	trimCmd.Flags().BoolP("right", "r", false, "only trim trailing characters")

	rootCmd.AddCommand(trimCmd)
}
