package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Git repository commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gitCmd).Standalone()

	rootCmd.AddCommand(gitCmd)
}
