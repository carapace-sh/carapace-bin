package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createIgnoreCmd = &cobra.Command{
	Use:   "create-ignore",
	Short: "Create a .gitignore per directory with SVN ignore properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createIgnoreCmd).Standalone()

	createIgnoreCmd.Flags().IntP("revision", "r", 0, "Refer to a specific revision")
	rootCmd.AddCommand(createIgnoreCmd)
}
