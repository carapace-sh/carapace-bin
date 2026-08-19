package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("deployment", "d", "", "Set the symbolic name of the new deployment")
	createCmd.Flags().StringP("include", "I", "", "Add path to Nix expression search path")

	createCmd.Flag("include").Nargs = 2

	rootCmd.AddCommand(createCmd)
}
