package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Cluster commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clusterCmd).Standalone()

	clusterCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(clusterCmd)
}
