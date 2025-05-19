package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List secure files for a project.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_listCmd).Standalone()

	securefile_listCmd.Flags().StringP("page", "p", "", "Page number.")
	securefile_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	securefileCmd.AddCommand(securefile_listCmd)
}
