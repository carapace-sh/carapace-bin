package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prepare your working directory for other commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("backend", "", "Configure the backend for this configuration.")
	initCmd.Flags().StringSlice("backend-config", []string{}, "Backend configuration")
	initCmd.Flags().Bool("force-copy", false, "Suppress prompts about copying state data.")
	initCmd.Flags().String("from-module", "", "Copy the contents of the given module into the target directory before initialization.")
	initCmd.Flags().String("get", "", "Download any modules for this configuration.")
	initCmd.Flags().Bool("ignore-remote-version", false, "A rare option used for the remote backend only.")
	initCmd.Flags().String("input", "", "Ask for input if necessary.")
	initCmd.Flags().String("lockfile", "", "Set a dependency lockfile mode.")
	initCmd.Flags().Bool("migrate-state", false, "Reconfigure the backend, and attempt to migrate any existing state.")
	initCmd.Flags().Bool("no-color", false, "If specified, output won't contain any color.")
	initCmd.Flags().Bool("plugin-dir", false, "Directory containing plugin binaries.")
	initCmd.Flags().Bool("reconfigure", false, "Reconfigure the backend, ignoring any saved configuration.")
	initCmd.Flags().String("upgrade", "", "ignore previously-downloaded objects and install the latest version allowed")
	rootCmd.AddCommand(initCmd)

	initCmd.Flag("backend").NoOptDefVal = " "
	initCmd.Flag("backend-config").NoOptDefVal = " "
	initCmd.Flag("from-module").NoOptDefVal = " "
	initCmd.Flag("get").NoOptDefVal = " "
	initCmd.Flag("input").NoOptDefVal = " "
	initCmd.Flag("lockfile").NoOptDefVal = " "
	initCmd.Flag("upgrade").NoOptDefVal = " "

	// TODO module completion
	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"backend":        action.ActionBool(),
		"backend-config": carapace.ActionFiles(),
		"get":            action.ActionBool(),
		"input":          action.ActionBool(),
		"lockfile":       carapace.ActionValues("readonly"),
		"upgrade":        action.ActionBool(),
	})
}
