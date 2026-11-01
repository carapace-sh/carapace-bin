package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "[experimental] List project daemons",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_lsCmd).Standalone()

	daemons_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	daemons_lsCmd.Flags().Bool("json", false, "Output in JSON format, including ports, URLs, and data directories")
	daemonsCmd.AddCommand(daemons_lsCmd)
}
