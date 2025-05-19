package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List log groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupsCmd).Standalone()

	groupsCmd.Flags().String("prefix", "", "log group prefix filter")
	rootCmd.AddCommand(groupsCmd)
}
