package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var mod_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "print module requirement graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_graphCmd).Standalone()
	mod_graphCmd.Flags().SetInterspersed(false)

	mod_graphCmd.Flags().StringS("go", "go", "", "report the module graph as loaded by their given Go version")
	mod_graphCmd.Flags().BoolS("x", "x", false, "print the commands graph executes")
	modCmd.AddCommand(mod_graphCmd)

	carapace.Gen(mod_graphCmd).FlagCompletion(carapace.ActionMap{
		"go": golang.ActionVersions(),
	})
}
