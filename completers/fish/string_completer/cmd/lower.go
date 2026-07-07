package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lowerCmd = &cobra.Command{
	Use:   "lower",
	Short: "Convert strings to lowercase",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lowerCmd).Standalone()

	lowerCmd.Flags().BoolP("quiet", "q", false, "suppress output")

	rootCmd.AddCommand(lowerCmd)
}
