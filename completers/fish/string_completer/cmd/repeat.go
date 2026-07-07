package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repeatCmd = &cobra.Command{
	Use:   "repeat",
	Short: "Repeat a string",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repeatCmd).Standalone()

	repeatCmd.Flags().BoolS("N", "N", false, "don't append a trailing newline")
	repeatCmd.Flags().StringP("count", "n", "", "number of repetitions")
	repeatCmd.Flags().StringP("max", "m", "", "maximum number of output characters")
	repeatCmd.Flags().Bool("no-newline", false, "don't append a trailing newline")
	repeatCmd.Flags().BoolP("quiet", "q", false, "suppress output")

	rootCmd.AddCommand(repeatCmd)
}
