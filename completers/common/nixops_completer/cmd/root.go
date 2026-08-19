package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
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

	rootCmd.Flags().Bool("confirm", false, "Confirm deployment")
	rootCmd.Flags().String("cores", "", "Set NIX_BUILD_CORES environment variable in builders")
	rootCmd.Flags().Bool("debug", false, "Turn on debugging output")
	rootCmd.Flags().StringP("deployment", "d", "", "Deployment to use")
	rootCmd.Flags().Bool("fallback", false, "Fall back on installation from source")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().Bool("keep-failed", false, "Keep temporary directories of failed builds")
	rootCmd.Flags().Bool("keep-going", false, "Keep going after failed builds")
	rootCmd.Flags().String("max-jobs", "", "Set maximum number of concurrent Nix builds")
	rootCmd.Flags().StringSlice("option", nil, "Set a Nix option")
	rootCmd.Flags().Bool("read-only-mode", false, "Run Nix evaluations in read-only mode")
	rootCmd.Flags().Bool("show-trace", false, "Print a Nix stack trace if evaluation fails")
	rootCmd.Flags().StringP("state", "s", "", "Path to state file")
	rootCmd.Flags().Bool("version", false, "Print NixOps's version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("list", "create", "modify", "clone", "delete", "info", "check", "set-args", "deploy", "send-keys", "destroy", "stop", "start", "reboot", "show-physical", "ssh", "ssh-for-each", "scp", "rename", "backup", "backup-status", "remove-backup", "clean-backups", "restore", "show-option", "list-generations", "rollback", "delete-generation", "show-console-output", "dump-nix-paths", "export", "import", "edit"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		nix.ActionPaths(),
	)
}
