package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clean_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Clean all profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clean_allCmd).Standalone()

	addCleanFlags(clean_allCmd)

	cleanCmd.AddCommand(clean_allCmd)
}
