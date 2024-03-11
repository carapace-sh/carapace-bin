package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:     "admin",
	Short:   "Operations requiring admin access on the Gitea instance",
	GroupID: "MISCELLANEOUS",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(adminCmd).Standalone()

	rootCmd.AddCommand(adminCmd)
}
