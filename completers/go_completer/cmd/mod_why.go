package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_whyCmd = &cobra.Command{
	Use:   "why",
	Short: "explain why packages or modules are needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_whyCmd).Standalone()

	mod_whyCmd.Flags().BoolS("m", "m", false, "treat arguments as a list of modules")
	mod_whyCmd.Flags().Bool("vendor", false, "exclude tests of dependencies")
	modCmd.AddCommand(mod_whyCmd)
}
