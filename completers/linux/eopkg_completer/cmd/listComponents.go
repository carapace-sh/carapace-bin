package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listComponentsCmd = &cobra.Command{
	Use:     "list-components",
	Aliases: []string{"lc"},
	Short:   "show all available components in the combined indexes of all installed repositories",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listComponentsCmd).Standalone()

	listComponentsCmd.Flags().BoolP("long", "l", false, "show full details on each component instead of just listing the names")
	listComponentsCmd.Flags().StringP("repository", "r", "", "only list components in the specified repository")

	rootCmd.AddCommand(listComponentsCmd)
}
