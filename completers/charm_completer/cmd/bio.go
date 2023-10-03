package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bioCmd = &cobra.Command{
	Use:    "bio",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bioCmd).Standalone()

	rootCmd.AddCommand(bioCmd)
}
