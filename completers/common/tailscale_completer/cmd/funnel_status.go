package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var funnel_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "View current funnel configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(funnel_statusCmd).Standalone()

	funnel_statusCmd.Flags().Bool("json", false, "output in JSON format")
	funnelCmd.AddCommand(funnel_statusCmd)
}
