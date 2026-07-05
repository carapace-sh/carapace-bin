package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computersystemCmd = &cobra.Command{
	Use:   "computersystem",
	Short: "computer system management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computersystemCmd).Standalone()
	rootCmd.AddCommand(computersystemCmd)
}
