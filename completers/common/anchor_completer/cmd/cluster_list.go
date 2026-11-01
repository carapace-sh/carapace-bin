package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints common cluster urls",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_listCmd).Standalone()

	cluster_listCmd.Flags().BoolP("help", "h", false, "Print help")
	clusterCmd.AddCommand(cluster_listCmd)
}
