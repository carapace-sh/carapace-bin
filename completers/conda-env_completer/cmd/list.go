package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the Conda environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	listCmd.Flags().Bool("json", false, "Report all output as json")
	listCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar")
	listCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace")
	rootCmd.AddCommand(listCmd)
}
