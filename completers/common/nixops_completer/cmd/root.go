package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nixops",
	Short: "NixOps deployment tool",
	Long:  "https://nixos.org/nixops/manual/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("confirm", false, "Confirm deployment")
	rootCmd.PersistentFlags().String("cores", "", "Set NIX_BUILD_CORES environment variable in builders")
	rootCmd.PersistentFlags().Bool("debug", false, "Turn on debugging output")
	rootCmd.PersistentFlags().StringP("deployment", "d", "", "Deployment to use")
	rootCmd.PersistentFlags().Bool("fallback", false, "Fall back on installation from source")
	rootCmd.PersistentFlags().Bool("help", false, "Show usage information")
	rootCmd.PersistentFlags().Bool("keep-failed", false, "Keep temporary directories of failed builds")
	rootCmd.PersistentFlags().Bool("keep-going", false, "Keep going after failed builds")
	rootCmd.PersistentFlags().String("max-jobs", "", "Set maximum number of concurrent Nix builds")
	rootCmd.PersistentFlags().StringSlice("option", nil, "Set a Nix option")
	rootCmd.PersistentFlags().Bool("read-only-mode", false, "Run Nix evaluations in read-only mode")
	rootCmd.PersistentFlags().Bool("show-trace", false, "Print a Nix stack trace if evaluation fails")
	rootCmd.PersistentFlags().StringP("state", "s", "", "Path to state file")
	rootCmd.PersistentFlags().Bool("version", false, "Print NixOps's version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("backup", "backup-status", "check", "clean-backups", "clone", "copy-closure", "create", "delete", "delete-generation", "delete-resources", "deploy", "destroy", "dump-nix-paths", "edit", "export", "import", "info", "list", "list-generations", "list-plugins", "modify", "mount", "reboot", "remove-backup", "rename", "restore", "rollback", "scp", "send-keys", "set-args", "show-arguments", "show-console-output", "show-option", "show-physical", "ssh", "ssh-for-each", "start", "stop", "unlock"),
	)
}
