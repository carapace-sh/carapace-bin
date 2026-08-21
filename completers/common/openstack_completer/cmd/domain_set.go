package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domain_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set domain properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domain_setCmd).Standalone()

	domain_setCmd.Flags().String("description", "", "New domain description")
	domain_setCmd.Flags().Bool("disable", false, "Disable domain")
	domain_setCmd.Flags().Bool("enable", false, "Enable domain")
	domain_setCmd.Flags().Bool("immutable", false, "Make resource immutable.")
	domain_setCmd.Flags().String("name", "", "New domain name")
	domain_setCmd.Flags().Bool("no-immutable", false, "Make resource mutable (default)")
	domainCmd.AddCommand(domain_setCmd)
}
