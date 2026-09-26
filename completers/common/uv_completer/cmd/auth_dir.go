package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Show the path to the uv credentials directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_dirCmd).Standalone()

	authCmd.AddCommand(auth_dirCmd)
}
