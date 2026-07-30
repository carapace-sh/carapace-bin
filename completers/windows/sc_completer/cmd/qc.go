package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qcCmd = &cobra.Command{
	Use:   "qc",
	Short: "query the configuration of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qcCmd).Standalone()
	rootCmd.AddCommand(qcCmd)
}
