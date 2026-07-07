package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lengthCmd = &cobra.Command{
	Use:   "length",
	Short: "Report string length",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lengthCmd).Standalone()

	lengthCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	lengthCmd.Flags().BoolP("visible", "V", false, "measure visible width")

	rootCmd.AddCommand(lengthCmd)
}
