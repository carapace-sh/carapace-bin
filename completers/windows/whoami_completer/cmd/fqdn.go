package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fqdnCmd = &cobra.Command{
	Use:   "fqdn",
	Short: "display fully qualified domain name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fqdnCmd).Standalone()
	rootCmd.AddCommand(fqdnCmd)
}
