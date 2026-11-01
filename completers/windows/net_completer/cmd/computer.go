package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var computerCmd = &cobra.Command{
	Use:   "computer",
	Short: "add or remove a computer from the domain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(computerCmd).Standalone()
	rootCmd.AddCommand(computerCmd)
}
