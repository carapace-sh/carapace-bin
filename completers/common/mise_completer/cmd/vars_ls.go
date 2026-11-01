package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vars_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List every var, with the file it comes from",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vars_lsCmd).Standalone()

	vars_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	vars_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	varsCmd.AddCommand(vars_lsCmd)
}
