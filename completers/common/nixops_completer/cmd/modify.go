package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "modify an existing deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modifyCmd).Standalone()

	modifyCmd.Flags().StringP("name", "n", "", "New symbolic name of deployment")

	rootCmd.AddCommand(modifyCmd)
}
