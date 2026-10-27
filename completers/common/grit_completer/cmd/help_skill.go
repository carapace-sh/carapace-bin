package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Print an agent skill (SKILL.md) that explains how to use grit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_skillCmd).Standalone()

	helpCmd.AddCommand(help_skillCmd)
}
