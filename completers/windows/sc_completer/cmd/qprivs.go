package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qprivsCmd = &cobra.Command{
	Use:   "qprivs",
	Short: "query the required privileges of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qprivsCmd).Standalone()
	rootCmd.AddCommand(qprivsCmd)
}
