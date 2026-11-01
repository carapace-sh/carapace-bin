package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qfailureCmd = &cobra.Command{
	Use:   "qfailure",
	Short: "query the failure actions of a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qfailureCmd).Standalone()
	rootCmd.AddCommand(qfailureCmd)
}
