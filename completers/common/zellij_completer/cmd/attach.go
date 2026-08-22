package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:     "attach",
	Short:   "Attach to a session",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().String("ca-cert", "", "Path to a custom CA certificate (PEM format) for verifying the remote server")
	attachCmd.Flags().BoolP("create", "c", false, "Create a session if one does not exist")
	attachCmd.Flags().BoolP("create-background", "b", false, "Create a detached session in the background if one does not exist")
	attachCmd.Flags().BoolP("force-run-commands", "f", false, "If resurrecting a dead session, immediately run all its commands on startup")
	attachCmd.Flags().Bool("forget", false, "Delete saved session before connecting")
	attachCmd.Flags().BoolP("help", "h", false, "Print help")
	attachCmd.Flags().String("index", "", "Number of the session index in the active sessions ordered creation date")
	attachCmd.Flags().Bool("insecure", false, "Skip TLS certificate validation (DANGEROUS — development only)")
	attachCmd.Flags().BoolP("remember", "r", false, "Save session for automatic re-authentication (4 weeks)")
	attachCmd.Flags().StringP("token", "t", "", "Authentication token for remote sessions")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"ca-cert": carapace.ActionFiles(),
	})

	carapace.Gen(attachCmd).PositionalCompletion(
		zellij.ActionSessions(),
	)
}
