package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qsidtypeCmd = &cobra.Command{
	Use:   "qsidtype",
	Short: "query the service SID type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qsidtypeCmd).Standalone()
	rootCmd.AddCommand(qsidtypeCmd)
}
