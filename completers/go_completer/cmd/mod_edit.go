package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var mod_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit go.mod from tools or scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_editCmd).Standalone()
	mod_editCmd.Flags().SetInterspersed(false)

	mod_editCmd.Flags().StringS("dropexclude", "dropexclude", "", "drop an exclusion")
	mod_editCmd.Flags().StringS("dropreplace", "dropreplace", "", "drop a module replacement")
	mod_editCmd.Flags().StringS("droprequire", "droprequire", "", "drop a requirement")
	mod_editCmd.Flags().StringS("exclude", "exclude", "", "add an exclusion")
	mod_editCmd.Flags().BoolS("fmt", "fmt", false, "reformats the go.mod file without making other changes")
	mod_editCmd.Flags().StringS("go", "go", "", "set the expected Go language version")
	mod_editCmd.Flags().BoolS("json", "json", false, "print the final go.mod file in JSON format instead instead of writing back")
	mod_editCmd.Flags().StringS("module", "module", "", "changes the module's path")
	mod_editCmd.Flags().BoolS("print", "print", false, "print the final go.mod in its text format instead of writing back")
	mod_editCmd.Flags().StringS("replace", "replace", "", "add a module replacement")
	mod_editCmd.Flags().StringS("require", "require", "", "add a requirement")
	modCmd.AddCommand(mod_editCmd)

	carapace.Gen(mod_editCmd).FlagCompletion(carapace.ActionMap{
		"dropexclude": golang.ActionModules(golang.ModuleOpts{Exclude: true, IncludeVersion: true}),
		"dropreplace": golang.ActionModules(golang.ModuleOpts{Replace: true}),
		"droprequire": golang.ActionModules(golang.ModuleOpts{Direct: true, IncludeVersion: false}),
		"exclude":     golang.ActionModuleSearch(),
		"go":          golang.ActionVersions(),
		"module":      carapace.ActionFiles(),
		"replace": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return golang.ActionModules(golang.ModuleOpts{Direct: true, Indirect: true}).Invoke(c).Suffix("=").ToA()
			default:
				return carapace.Batch(
					carapace.ActionDirectories().ChdirF(traverse.Parent("go.mod")),
					golang.ActionModuleSearch().UnlessF(condition.CompletingPath),
				).ToA()
			}
		}),
		"require": golang.ActionModuleSearch(),
	})
}
