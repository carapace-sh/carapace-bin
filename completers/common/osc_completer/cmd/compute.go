package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Compute service commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computeCmd).Standalone()
	rootCmd.AddCommand(computeCmd)

	subresources := []struct {
		use   string
		short string
	}{
		{"server", "Server commands"},
		{"flavor", "Flavor commands"},
		{"keypair", "Keypair commands"},
		{"aggregate", "Host Aggregate commands"},
		{"availability-zone", "Availability Zone commands"},
		{"hypervisor", "Hypervisor commands"},
		{"quota-set", "Quota Set commands"},
		{"server-group", "Server Group commands"},
	}

	for _, r := range subresources {
		sub := &cobra.Command{Use: r.use, Short: r.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		computeCmd.AddCommand(sub)
		addCrudSubcommands(sub)
	}
}

func addCrudSubcommands(parent *cobra.Command) {
	cmds := []struct {
		use   string
		short string
	}{
		{"list", "List resources"},
		{"show", "Show resource details"},
		{"create", "Create a resource"},
		{"delete", "Delete a resource"},
		{"set", "Update a resource"},
	}
	for _, c := range cmds {
		sub := &cobra.Command{Use: c.use, Short: c.short, Run: func(cmd *cobra.Command, args []string) {}}
		carapace.Gen(sub).Standalone()
		parent.AddCommand(sub)
	}
}
