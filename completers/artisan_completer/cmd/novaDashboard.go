package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaDashboardCmd = &cobra.Command{
	Use:   "nova:dashboard <name>",
	Short: "Create a new dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaDashboardCmd).Standalone()

	rootCmd.AddCommand(novaDashboardCmd)
}
