package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var privsCmd = &cobra.Command{
	Use:   "privs",
	Short: "change the required privileges of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(privsCmd).Standalone()
	rootCmd.AddCommand(privsCmd)
}
