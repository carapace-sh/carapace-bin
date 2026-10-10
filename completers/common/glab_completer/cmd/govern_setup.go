package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var govern_setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure this machine to record AI agent sessions. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(govern_setupCmd).Standalone()

	govern_setupCmd.Flags().StringSlice("agents", nil, "Also sync sessions from these agents, found by the fallback periodic sync without hooks: codex, cursor. Replaces the agents enabled by an earlier run. Multiple agents can be comma-separated or specified by repeating the flag.")
	govern_setupCmd.Flags().Bool("no-fallback-sync", false, "Install only the hooks, without the fallback periodic sync job.")
	govern_setupCmd.Flags().Bool("uninstall", false, "Remove the fallback periodic sync job and stop syncing Codex and Cursor sessions. The hooks are left in place.")
	govern_setupCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt.")
	governCmd.AddCommand(govern_setupCmd)

	carapace.Gen(govern_setupCmd).FlagCompletion(carapace.ActionMap{
		"agents": carapace.ActionValues("codex", "cursor"),
	})
}
