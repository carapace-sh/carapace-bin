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
	rootCmd.Flags().Bool("debug", false, "Turn on debugging output")
	rootCmd.Flags().StringP("deployment", "d", "", "Deployment to use")
	rootCmd.Flags().Bool("help", false, "Show usage information")
	rootCmd.Flags().StringP("state", "s", "", "Path to state file")

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
