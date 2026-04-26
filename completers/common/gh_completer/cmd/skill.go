package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillCmd = &cobra.Command{
	Use:     "skill <command>",
	Short:   "Install and manage agent skills (preview)",
	GroupID: "core",
	Aliases: []string{"skills"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillCmd).Standalone()

	rootCmd.AddCommand(skillCmd)
}
