package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var privCmd = &cobra.Command{
	Use:   "priv",
	Short: "display privileges",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(privCmd).Standalone()
	rootCmd.AddCommand(privCmd)
}
