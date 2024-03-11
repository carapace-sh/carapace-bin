package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actionsCmd = &cobra.Command{
	Use:    "actions",
	Short:  "Learn about working with GitHub Actions",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actionsCmd).Standalone()

	rootCmd.AddCommand(actionsCmd)
}
