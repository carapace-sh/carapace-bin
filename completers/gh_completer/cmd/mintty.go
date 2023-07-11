package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var minttyCmd = &cobra.Command{
	Use:    "mintty",
	Short:  "Information about using gh with MinTTY",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(minttyCmd).Standalone()

	rootCmd.AddCommand(minttyCmd)
}
