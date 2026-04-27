package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var objectStoreCmd = &cobra.Command{
	Use:   "object-store",
	Short: "Object Store service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(objectStoreCmd).Standalone()
	rootCmd.AddCommand(objectStoreCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"account", "Account commands"},
		{"container", "Container commands"},
		{"object", "Object commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		objectStoreCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
