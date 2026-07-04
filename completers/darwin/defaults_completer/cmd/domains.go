package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "List all domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domainsCmd).Standalone()
	rootCmd.AddCommand(domainsCmd)
}
