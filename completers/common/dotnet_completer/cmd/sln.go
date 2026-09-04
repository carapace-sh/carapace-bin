package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var slnCmd = &cobra.Command{
	Use:   "sln",
	Short: "modify solutions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(slnCmd).Standalone()
	rootCmd.AddCommand(slnCmd)
}
