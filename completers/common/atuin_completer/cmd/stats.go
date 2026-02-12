package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Calculate statistics for your history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().StringP("count", "c", "", "How many top commands to list")
	statsCmd.Flags().BoolP("help", "h", false, "Print help")
	statsCmd.Flags().StringP("ngram-size", "n", "", "The number of consecutive commands to consider")
	rootCmd.AddCommand(statsCmd)
}
