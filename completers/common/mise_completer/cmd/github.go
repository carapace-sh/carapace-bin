package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:    "github",
	Short:  "GitHub commands (deprecated; use `mise token github`)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(githubCmd).Standalone()

	githubCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(githubCmd)
}
