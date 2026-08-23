package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zprint",
	Short: "show information about kernel zones",
	Long:  "https://keith.github.io/xcode-manpages/zprint.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("column", "c", false, "Print zone info in columns")
	rootCmd.Flags().BoolP("delta", "d", false, "Display deltas over time")
	rootCmd.Flags().BoolP("heading", "h", false, "Show headings for the columns")
	rootCmd.Flags().BoolP("no-wired", "L", false, "Do not show wired memory information")
	rootCmd.Flags().BoolP("sort", "s", false, "Sort zones showing the zone wasting the most memory first")
	rootCmd.Flags().BoolP("total", "t", false, "Calculate total size of allocations over the life of the zone")
	rootCmd.Flags().BoolP("waste", "w", false, "Calculate how much space is allocated but not in use")
	rootCmd.Flags().BoolP("wired", "l", false, "Show all wired memory information")
}