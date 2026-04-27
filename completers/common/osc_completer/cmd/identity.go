package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identityCmd = &cobra.Command{
	Use:   "identity",
	Short: "Identity (Keystone) service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identityCmd).Standalone()
	rootCmd.AddCommand(identityCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"project", "Project commands"},
		{"user", "User commands"},
		{"group", "Group commands"},
		{"role", "Role commands"},
		{"domain", "Domain commands"},
		{"region", "Region commands"},
		{"application-credential", "Application Credential commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		identityCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
