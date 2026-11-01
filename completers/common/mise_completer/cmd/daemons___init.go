package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons___initCmd = &cobra.Command{
	Use:    "__init",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons___initCmd).Standalone()

	daemons___initCmd.Flags().String("context", "", "Resolved preset ports and options as a JSON object.")
	daemons___initCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons___initCmd)
}
