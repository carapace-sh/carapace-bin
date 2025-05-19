package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "workspace maintenance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workCmd).Standalone()
	workCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(workCmd)
}
