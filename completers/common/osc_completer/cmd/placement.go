package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var placementCmd = &cobra.Command{
	Use:   "placement",
	Short: "Placement service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(placementCmd).Standalone()
	rootCmd.AddCommand(placementCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"resource-provider", "Resource Provider commands"},
		{"resource-class", "Resource Class commands"},
		{"allocation", "Allocation commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		placementCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
