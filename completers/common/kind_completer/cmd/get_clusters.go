package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var get_clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Lists existing kind clusters by their name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_clustersCmd).Standalone()

	getCmd.AddCommand(get_clustersCmd)
}
