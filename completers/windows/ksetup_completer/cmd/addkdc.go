package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addkdcCmd = &cobra.Command{
	Use:   "addkdc",
	Short: "add a KDC address for the realm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addkdcCmd).Standalone()
	rootCmd.AddCommand(addkdcCmd)
}
