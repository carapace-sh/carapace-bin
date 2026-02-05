package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query [KEYWORDS]...",
	Short: "Search for a directory in the database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	queryCmd.Flags().BoolP("all", "a", false, "Show unavailable directories")
	queryCmd.Flags().String("base-dir", "", "Only search within this directory")
	queryCmd.Flags().String("exclude", "", "Exclude the current directory")
	queryCmd.Flags().BoolP("help", "h", false, "Print help")
	queryCmd.Flags().BoolP("interactive", "i", false, "Use interactive selection")
	queryCmd.Flags().BoolP("list", "l", false, "List all matching directories")
	queryCmd.Flags().BoolP("score", "s", false, "Print score with results")
	queryCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(queryCmd)

	carapace.Gen(queryCmd).FlagCompletion(carapace.ActionMap{
		"base-dir": carapace.ActionDirectories(),
		"exclude":  carapace.ActionDirectories(),
	})
}
