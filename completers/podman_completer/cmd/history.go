package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history [options] IMAGE",
	Short: "Show history of a specified image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().String("format", "", "Change the output to JSON or a Go template")
	historyCmd.Flags().BoolP("human", "H", false, "Display sizes and dates in human readable format")
	historyCmd.Flags().Bool("no-trunc", false, "Do not truncate the output")
	historyCmd.Flags().BoolP("quiet", "q", false, "Display the numeric IDs only")
	rootCmd.AddCommand(historyCmd)
}
