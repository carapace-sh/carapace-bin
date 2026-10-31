package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Lists existing kind nodes by their name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_nodesCmd).Standalone()

	get_nodesCmd.Flags().BoolP("all-clusters", "A", false, "If present, list all the available nodes across all cluster contexts. Current context is ignored even if specified with --name.")
	get_nodesCmd.Flags().StringP("name", "n", "kind", "the cluster context name")

	carapace.Gen(get_nodesCmd).FlagCompletion(carapace.ActionMap{
		"name": action.ActionClusterNames(),
	})

	getCmd.AddCommand(get_nodesCmd)
}
