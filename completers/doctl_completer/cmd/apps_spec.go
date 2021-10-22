package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_specCmd = &cobra.Command{
	Use:   "spec",
	Short: "Display commands for working with app specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_specCmd).Standalone()
	appsCmd.AddCommand(apps_specCmd)
}
