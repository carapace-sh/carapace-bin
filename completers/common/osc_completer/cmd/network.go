package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Network service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkCmd).Standalone()
	rootCmd.AddCommand(networkCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"network", "Network commands"},
		{"subnet", "Subnet commands"},
		{"port", "Port commands"},
		{"router", "Router commands"},
		{"floating-ip", "Floating IP commands"},
		{"security-group", "Security Group commands"},
		{"security-group-rule", "Security Group Rule commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		networkCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
