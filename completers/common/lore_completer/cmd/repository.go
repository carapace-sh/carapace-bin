package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Repository commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repositoryCmd).Standalone()

	repositoryCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(repositoryCmd)
}
