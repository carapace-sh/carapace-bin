package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attestationCmd = &cobra.Command{
	Use:     "attestation [subcommand]",
	Short:   "Work with artifact attestations",
	Aliases: []string{"at"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestationCmd).Standalone()

	rootCmd.AddCommand(attestationCmd)
}
