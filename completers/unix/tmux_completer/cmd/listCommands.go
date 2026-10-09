package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCommandsCmd = &cobra.Command{
	Use:     "list-commands",
	Aliases: []string{"lscm"},
	Short:   "list supported sub-commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCommandsCmd).Standalone()

	listCommandsCmd.Flags().StringS("F", "F", "", "specify format")
	rootCmd.AddCommand(listCommandsCmd)
}
