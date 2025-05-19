package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var mod_tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "add missing and remove unused modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_tidyCmd).Standalone()
	mod_tidyCmd.Flags().SetInterspersed(false)

	mod_tidyCmd.Flags().StringS("compat", "compat", "", "preserve additional checksums needed for given version")
	mod_tidyCmd.Flags().BoolS("e", "e", false, "attempt to proceed despite errors")
	mod_tidyCmd.Flags().StringS("go", "go", "", "update the 'go' directive to given version")
	mod_tidyCmd.Flags().BoolS("v", "v", false, "print information about removed modules")
	modCmd.AddCommand(mod_tidyCmd)

	carapace.Gen(mod_tidyCmd).FlagCompletion(carapace.ActionMap{
		"compat": golang.ActionVersions(),
		"go":     golang.ActionVersions(),
	})
}
