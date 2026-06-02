package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clean_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Clean the current user's profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clean_userCmd).Standalone()

	addCleanFlags(clean_userCmd)

	cleanCmd.AddCommand(clean_userCmd)
}
