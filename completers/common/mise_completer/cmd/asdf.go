// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var asdfCmd = &cobra.Command{
	Use:    "asdf",
	Short:  "[internal] simulates asdf for plugins that call \"asdf\" internally",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(asdfCmd).Standalone()

	asdfCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(asdfCmd)
}
