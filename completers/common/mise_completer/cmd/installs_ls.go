package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var installs_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List the installations of the identity install layout",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installs_lsCmd).Standalone()

	installs_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	installs_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	installs_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	installsCmd.AddCommand(installs_lsCmd)

	carapace.Gen(installs_lsCmd).PositionalAnyCompletion(
		mise.ActionTools(),
	)
}
