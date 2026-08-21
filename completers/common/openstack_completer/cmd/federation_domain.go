package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_domainCmd).Standalone()

	federationCmd.AddCommand(federation_domainCmd)
}
