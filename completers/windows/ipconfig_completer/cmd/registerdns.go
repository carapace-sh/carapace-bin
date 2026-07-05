package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registerdnsCmd = &cobra.Command{
	Use:   "registerdns",
	Short: "initiate manual dynamic registration of DNS names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registerdnsCmd).Standalone()
	rootCmd.AddCommand(registerdnsCmd)
}
