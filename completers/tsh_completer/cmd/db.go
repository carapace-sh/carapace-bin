package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "View and control proxied databases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbCmd).Standalone()

	dbCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	rootCmd.AddCommand(dbCmd)

	carapace.Gen(dbCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
	})
}
