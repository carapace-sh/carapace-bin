package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerInfraCmd = &cobra.Command{
	Use:     "container-infrastructure",
	Short:   "Container Infrastructure Management service commands",
	Aliases: []string{"container-infrastructure-management", "container"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerInfraCmd).Standalone()
	rootCmd.AddCommand(containerInfraCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"cluster", "Cluster commands"},
		{"cluster-template", "Cluster Template commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		containerInfraCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}
