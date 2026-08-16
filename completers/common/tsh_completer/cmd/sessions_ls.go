package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sessions_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List active sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessions_lsCmd).Standalone()

	sessions_lsCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	sessions_lsCmd.Flags().String("kind", "ssh,k8s,db,app,desktop", "Filter by session kind(s).")
	sessionsCmd.AddCommand(sessions_lsCmd)
}
