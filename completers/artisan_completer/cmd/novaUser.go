package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaUserCmd = &cobra.Command{
	Use:   "nova:user",
	Short: "Create a new user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaUserCmd).Standalone()

	rootCmd.AddCommand(novaUserCmd)
}
