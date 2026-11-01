package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delkdcCmd = &cobra.Command{
	Use:   "delkdc",
	Short: "delete a KDC address for the realm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delkdcCmd).Standalone()
	rootCmd.AddCommand(delkdcCmd)
}
