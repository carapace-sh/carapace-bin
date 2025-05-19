package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var create_secretCmd = &cobra.Command{
	Use:   "secret (docker-registry | generic | tls)",
	Short: "Create a secret using a specified subcommand",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secretCmd).Standalone()

	createCmd.AddCommand(create_secretCmd)
}
