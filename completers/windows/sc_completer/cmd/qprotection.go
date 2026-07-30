package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qprotectionCmd = &cobra.Command{
	Use:   "qprotection",
	Short: "query the process protection level of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qprotectionCmd).Standalone()
	rootCmd.AddCommand(qprotectionCmd)
}
