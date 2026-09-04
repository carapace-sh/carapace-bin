package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsCmd = &cobra.Command{
	Use:     "apps",
	Short:   "View and control proxied applications.",
	Aliases: []string{"app"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsCmd).Standalone()

	appsCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	rootCmd.AddCommand(appsCmd)
}
