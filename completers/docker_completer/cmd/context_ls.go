package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_lsCmd = &cobra.Command{
	Use:     "ls [OPTIONS]",
	Short:   "List contexts",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_lsCmd).Standalone()

	context_lsCmd.Flags().String("format", "", "Format output using a custom template:")
	context_lsCmd.Flags().BoolP("quiet", "q", false, "Only show context names")
	contextCmd.AddCommand(context_lsCmd)
}
