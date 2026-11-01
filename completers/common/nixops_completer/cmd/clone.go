package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone an existing deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().StringP("name", "n", "", "Symbolic name of the cloned deployment")

	rootCmd.AddCommand(cloneCmd)
}
