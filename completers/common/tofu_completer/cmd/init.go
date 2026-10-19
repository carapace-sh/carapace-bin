package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [options]",
	Short:   "Prepare your working directory for other commands",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolS("backend", "backend", false, "Disable backend or cloud backend initialization.")
	initCmd.Flags().StringSliceS("backend-config", "backend-config", nil, "Configuration to be merged with what is in the configuration file's 'backend' block.")
	initCmd.Flags().BoolS("cloud", "cloud", false, "Disable backend or cloud backend initialization.")
	initCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	initCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	initCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	initCmd.Flags().BoolS("force-copy", "force-copy", false, "Suppress prompts about copying state data.")
	initCmd.Flags().StringS("from-module", "from-module", "", "Copy the contents of the given module into the target directory before initialization.")
	initCmd.Flags().BoolS("get", "get", false, "Disable downloading modules for this configuration.")
	initCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	initCmd.Flags().BoolS("input", "input", false, "Disable interactive prompts.")
	initCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	initCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	initCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during backend migration.")
	initCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	initCmd.Flags().StringS("lockfile", "lockfile", "", "Set a dependency lockfile mode.")
	initCmd.Flags().BoolS("migrate-state", "migrate-state", false, "Reconfigure a backend, and attempt to migrate any existing state.")
	initCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	initCmd.Flags().StringArrayS("plugin-dir", "plugin-dir", nil, "Directory containing plugin binaries. This overrides all other plugin search directory settings.")
	initCmd.Flags().BoolS("reconfigure", "reconfigure", false, "Reconfigure a backend, ignoring any saved configuration.")
	initCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	initCmd.Flags().BoolS("upgrade", "upgrade", false, "Install the latest module and provider versions allowed within configured constraints.")
	initCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	initCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(initCmd)

	initCmd.Flag("backend-config").NoOptDefVal = " "
	initCmd.Flag("from-module").NoOptDefVal = " "
	initCmd.Flag("json-into").NoOptDefVal = " "
	initCmd.Flag("lock-timeout").NoOptDefVal = " "
	initCmd.Flag("lockfile").NoOptDefVal = " "
	initCmd.Flag("plugin-dir").NoOptDefVal = " "
	initCmd.Flag("test-directory").NoOptDefVal = " "
	initCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"backend-config": carapace.ActionFiles(),
		"from-module":    carapace.ActionDirectories(),
		"json-into":      carapace.ActionFiles(),
		"lockfile":       carapace.ActionValues("readonly"),
		"plugin-dir":     carapace.ActionDirectories(),
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})
}
