package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/go_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mod_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit go.mod from tools or scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_editCmd).Standalone()

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

	// TODO complete more flags
	carapace.Gen(mod_editCmd).FlagCompletion(carapace.ActionMap{
		"droprequire": action.ActionModules(false),
		"module":      carapace.ActionFiles(),
		"require": carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				opts := git.SearchOpts{}.Default()
				opts.Prefix = false
				return git.ActionRepositorySearch(opts)
			case 1:
				return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + c.Parts[0], Tags: true})
			default:
				return carapace.ActionValues()
			}
		}),
	})
}
