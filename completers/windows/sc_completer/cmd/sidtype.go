package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sidtypeCmd = &cobra.Command{
	Use:   "sidtype",
	Short: "change the service SID type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sidtypeCmd).Standalone()
	rootCmd.AddCommand(sidtypeCmd)
}
