package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a metadata remote source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_addCmd).Standalone()

	remote_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	remote_addCmd.Flags().Bool("init", false, "Initialize the remote with a README commit on `refs/{namespace}/main` when no metadata refs exist there yet")
	remote_addCmd.Flags().String("name", "meta", "Remote name (default: meta)")
	remote_addCmd.Flags().String("namespace", "", "Metadata namespace to use (default: from git config or \"meta\")")
	remoteCmd.AddCommand(remote_addCmd)
}
