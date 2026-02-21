package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "DNS service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dnsCmd).Standalone()
	rootCmd.AddCommand(dnsCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"zone", "Zone commands"},
		{"recordset", "Recordset commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		dnsCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
