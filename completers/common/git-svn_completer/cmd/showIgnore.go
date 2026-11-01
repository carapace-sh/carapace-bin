package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showIgnoreCmd = &cobra.Command{
	Use:   "show-ignore",
	Short: "Show .gitignore patterns from SVN ignore properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showIgnoreCmd).Standalone()

	showIgnoreCmd.Flags().IntP("revision", "r", 0, "Refer to a specific revision")
	rootCmd.AddCommand(showIgnoreCmd)
}
