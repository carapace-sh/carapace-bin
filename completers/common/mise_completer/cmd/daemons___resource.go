// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons___resourceCmd = &cobra.Command{
	Use:    "__resource",
	Short:  "A project-owned readiness process: it never starts or stops the server itself.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons___resourceCmd).Standalone()

	daemons___resourceCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons___resourceCmd)
}
