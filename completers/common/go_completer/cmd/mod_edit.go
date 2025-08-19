package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var mod_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit go.mod from tools or scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_editCmd).Standalone()
	mod_editCmd.Flags().SetInterspersed(false)

	mod_editCmd.Flags().StringS("C", "C", "", "Change to dir before running the command")
	mod_editCmd.Flags().StringArrayS("dropexclude", "dropexclude", nil, "drop an exclusion")
	mod_editCmd.Flags().StringArrayS("dropgodebug", "dropgodebug", nil, "drop any existing godebug lines")
	mod_editCmd.Flags().StringArrayS("dropignore", "dropignore", nil, "drop a ignore declaration")
	mod_editCmd.Flags().StringArrayS("dropreplace", "dropreplace", nil, "drop a module replacement")
	mod_editCmd.Flags().StringArrayS("droprequire", "droprequire", nil, "drop a requirement")
	mod_editCmd.Flags().StringArrayS("dropretract", "dropretract", nil, "drop a retraction")
	mod_editCmd.Flags().StringArrayS("droptool", "droptool", nil, "drop a tool declaration")
	mod_editCmd.Flags().StringArrayS("exclude", "exclude", nil, "add an exclusion")
	mod_editCmd.Flags().BoolS("fmt", "fmt", false, "reformats the go.mod file without making other changes")
	mod_editCmd.Flags().StringS("go", "go", "", "set the expected Go language version")
	mod_editCmd.Flags().StringArrayS("godebug", "godebug", nil, "add a godebug key=value line")
	mod_editCmd.Flags().StringArrayS("ignore", "ignore", nil, "add a ignore declaration")
	mod_editCmd.Flags().BoolS("json", "json", false, "print the final go.mod file in JSON format instead instead of writing back")
	mod_editCmd.Flags().StringS("module", "module", "", "changes the module's path")
	mod_editCmd.Flags().BoolS("n", "n", false, "print the commands but do not run them")
	mod_editCmd.Flags().BoolS("print", "print", false, "print the final go.mod in its text format instead of writing back")
	mod_editCmd.Flags().StringArrayS("replace", "replace", nil, "add a module replacement")
	mod_editCmd.Flags().StringArrayS("require", "require", nil, "add a requirement")
	mod_editCmd.Flags().StringArrayS("retract", "retract", nil, "add a retraction")
	mod_editCmd.Flags().StringArrayS("tool", "tool", nil, "add a tool declaration")
	mod_editCmd.Flags().StringS("toolchain", "toolchain", "", "set the Go toolchain to use")
	mod_editCmd.Flags().BoolS("x", "x", false, "print the commands")
	modCmd.AddCommand(mod_editCmd)

	carapace.Gen(mod_editCmd).FlagCompletion(carapace.ActionMap{
		"C":           carapace.ActionDirectories(),
		"dropexclude": golang.ActionModules(golang.ModuleOpts{Exclude: true, IncludeVersion: true}),
		"dropgodebug": golang.ActionModGodebugs(),
		"dropignore":  golang.ActionModIgnores(),
		"dropreplace": golang.ActionModules(golang.ModuleOpts{Replace: true}),
		"droprequire": golang.ActionModules(golang.ModuleOpts{Direct: true, IncludeVersion: false}),
		"dropretract": golang.ActionModRetracts(),
		"droptool":    golang.ActionModTools(),
		"exclude":     golang.ActionModuleSearch(),
		"go":          golang.ActionVersions(),
		"godebug":     golang.ActionGodebugKeyValues(),
		"ignore": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "./") {
				return carapace.ActionFiles().ChdirF(traverse.Parent("go.mod")) // TODO add traverse for this
			}
			return carapace.ActionFiles()
		}),
		"module": git.ActionRepositorySearch(git.SearchOpts{Prefix: false, Github: true, Gitlab: true}),
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
		"retract": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "[") {
				return carapace.ActionMultiPartsN(",", 2, func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return golang.ActionModVersions().Prefix("[").Suffix(",")
					default:
						return golang.ActionModVersions().Suffix("]")
					}
				})
			}
			return golang.ActionModVersions()
		}),
		"tool":      git.ActionRepositorySearch(git.SearchOpts{Prefix: false, Github: true, Gitlab: true}),
		"toolchain": golang.ActionVersions().Prefix("go"),
	})

	carapace.Gen(mod_editCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("C").Value.String())
	})
}
