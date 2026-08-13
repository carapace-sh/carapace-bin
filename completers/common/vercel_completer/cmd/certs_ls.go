package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certs_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Show all available certificates",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certs_lsCmd).Standalone()

	certs_lsCmd.Flags().String("limit", "", "Number of results per page")
	certs_lsCmd.Flags().String("next", "", "Show next page of results")

	certsCmd.AddCommand(certs_lsCmd)
}
