package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listComponentsCmd = &cobra.Command{
	Use:     "list-components",
	Short:   "List available components",
	Aliases: []string{"lc"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listComponentsCmd).Standalone()

	listComponentsCmd.Flags().BoolP("long", "l", false, "Show in long format")
	listComponentsCmd.Flags().StringP("repository", "r", "", "Name of the repository")

	rootCmd.AddCommand(listComponentsCmd)
}
