package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var securefile_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List secure files in a project.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_listCmd).Standalone()

	securefile_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	securefile_listCmd.Flags().StringP("page", "p", "", "Page number.")
	securefile_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page.")
	securefileCmd.AddCommand(securefile_listCmd)

	carapace.Gen(securefile_listCmd).FlagCompletion(carapace.ActionMap{
		"jq": jq.ActionFilters(),
	})
}
