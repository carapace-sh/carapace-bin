package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit go.mod from tools or scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_editCmd).Standalone()

	mod_editCmd.Flags().String("dropexclude", "", "drop an exclusion")
	mod_editCmd.Flags().String("dropreplace", "", "drop a module replacement")
	mod_editCmd.Flags().String("droprequire", "", "drop a requirement")
	mod_editCmd.Flags().String("exclude", "", "add an exclusion")
	mod_editCmd.Flags().Bool("fmt", false, "reformats the go.mod file without making other changes")
	mod_editCmd.Flags().String("go", "", "set the expected Go language version")
	mod_editCmd.Flags().Bool("json", false, "print the final go.mod file in JSON format instead instead of writing back")
	mod_editCmd.Flags().String("module", "", "changes the module's path")
	mod_editCmd.Flags().Bool("print", false, "print the final go.mod in its text format instead of writing back")
	mod_editCmd.Flags().String("replace", "", "add a module replacement")
	mod_editCmd.Flags().String("require", "", "add a requirement")
	modCmd.AddCommand(mod_editCmd)

	carapace.Gen(mod_editCmd).FlagCompletion(carapace.ActionMap{
		"module": carapace.ActionFiles(),
	})
}
